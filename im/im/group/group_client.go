package group

import (
	"imservice/im"
	"time"
)
import "sync/atomic"
import "errors"
import log "github.com/sirupsen/logrus"
import "imservice/storage"

type GroupClient struct {
	*im.Connection
}

func (client *GroupClient) HandleSuperGroupMessage(msg *im.IMMessage, group *Group) (int64, int64, error) {
	m := &im.Message{cmd: im.MSG_GROUP_IM, version: im.DEFAULT_VERSION, body: msg}
	msgid, prev_msgid, err := im.rpc_storage.SaveGroupMessage(client.appid, msg.receiver, client.device_ID, m)
	if err != nil {
		log.Errorf("save group message:%d %d err:%s", msg.sender, msg.receiver, err)
		return 0, 0, err
	}

	//推送外部通知
	im.PushGroupMessage(client.appid, group, m)

	m.meta = &im.Metadata{sync_key: msgid, prev_sync_key: prev_msgid}
	m.flag = im.MESSAGE_FLAG_PUSH | im.MESSAGE_FLAG_SUPER_GROUP
	client.SendGroupMessage(group, m)

	notify := &im.Message{cmd: im.MSG_SYNC_GROUP_NOTIFY, body: &im.GroupSyncKey{group_id: msg.receiver, sync_key: msgid}}
	client.SendGroupMessage(group, notify)

	return msgid, prev_msgid, nil
}

func (client *GroupClient) HandleGroupMessage(im *im.IMMessage, group *Group) (int64, int64, error) {
	gm := &im.PendingGroupMessage{}
	gm.appid = client.appid
	gm.sender = im.sender
	gm.device_ID = client.device_ID
	gm.gid = im.receiver
	gm.timestamp = im.timestamp

	members := group.Members()
	gm.members = make([]int64, len(members))
	i := 0
	for uid := range members {
		gm.members[i] = uid
		i += 1
	}

	gm.content = im.content
	deliver := im.GetGroupMessageDeliver(group.gid)
	m := &im.Message{cmd: im.MSG_PENDING_GROUP_MESSAGE, body: gm}

	c := make(chan *im.Metadata, 1)
	callback_id := deliver.SaveMessage(m, c)
	defer deliver.RemoveCallback(callback_id)
	select {
	case meta := <-c:
		return meta.sync_key, meta.prev_sync_key, nil
	case <-time.After(time.Second * 2):
		log.Errorf("save group message:%d %d timeout", im.sender, im.receiver)
		return 0, 0, errors.New("timeout")
	}
}

func (client *GroupClient) HandleGroupIMMessage(message *im.Message) {
	msg := message.body.(*im.IMMessage)
	seq := message.seq
	if client.uid == 0 {
		log.Warning("client has't been authenticated")
		return
	}

	if msg.sender != client.uid {
		log.Warningf("im message sender:%d client uid:%d\n", msg.sender, client.uid)
		return
	}
	if message.flag&im.MESSAGE_FLAG_TEXT != 0 {
		im.FilterDirtyWord(msg)
	}

	msg.timestamp = int32(time.Now().Unix())

	deliver := im.GetGroupMessageDeliver(msg.receiver)
	group := deliver.LoadGroup(msg.receiver)
	if group == nil {
		ack := &im.Message{cmd: im.MSG_ACK, body: &im.MessageACK{seq: int32(seq), status: im.ACK_GROUP_NONEXIST}}
		client.EnqueueMessage(ack)
		log.Warning("can't find group:", msg.receiver)
		return
	}

	if !group.IsMember(msg.sender) {
		ack := &im.Message{cmd: im.MSG_ACK, body: &im.MessageACK{seq: int32(seq), status: im.ACK_NOT_GROUP_MEMBER}}
		client.EnqueueMessage(ack)
		log.Warningf("sender:%d is not group member", msg.sender)
		return
	}

	if group.GetMemberMute(msg.sender) {
		log.Warningf("sender:%d is mute in group", msg.sender)
		return
	}

	var meta *im.Metadata
	var flag int
	if group.super {
		msgid, prev_msgid, err := client.HandleSuperGroupMessage(msg, group)
		if err == nil {
			meta = &im.Metadata{sync_key: msgid, prev_sync_key: prev_msgid}
		}
		flag = im.MESSAGE_FLAG_SUPER_GROUP
	} else {
		msgid, prev_msgid, err := client.HandleGroupMessage(msg, group)
		if err == nil {
			meta = &im.Metadata{sync_key: msgid, prev_sync_key: prev_msgid}
		}
	}

	ack := &im.Message{cmd: im.MSG_ACK, flag: flag, body: &im.MessageACK{seq: int32(seq)}, meta: meta}
	r := client.EnqueueMessage(ack)
	if !r {
		log.Warning("send group message ack error")
	}

	atomic.AddInt64(&im.server_summary.in_message_count, 1)
	log.Infof("group message sender:%d group id:%d super:%v", msg.sender, msg.receiver, group.super)
	if meta != nil {
		log.Info("group message ack meta:", meta.sync_key, meta.prev_sync_key)
	}
}

func (client *GroupClient) HandleGroupSync(group_sync_key *im.GroupSyncKey) {
	if client.uid == 0 {
		return
	}

	group_id := group_sync_key.group_id

	group := im.group_manager.FindGroup(group_id)
	if group == nil {
		log.Warning("can't find group:", group_id)
		return
	}

	if !group.IsMember(client.uid) {
		log.Warningf("sender:%d is not group member", client.uid)
		return
	}

	ts := group.GetMemberTimestamp(client.uid)

	last_id := group_sync_key.sync_key
	if last_id == 0 {
		last_id = im.GetGroupSyncKey(client.appid, client.uid, group_id)
	}

	log.Info("sync group message...", group_sync_key.sync_key, last_id)
	gh, err := im.rpc_storage.SyncGroupMessage(client.appid, client.uid, client.device_ID, group_sync_key.group_id, last_id, int32(ts))
	if err != nil {
		log.Warning("sync message err:", err)
		return
	}
	messages := gh.Messages

	sk := &im.GroupSyncKey{sync_key: last_id, group_id: group_id}
	client.EnqueueMessage(&im.Message{cmd: im.MSG_SYNC_GROUP_BEGIN, body: sk})
	for i := len(messages) - 1; i >= 0; i-- {
		msg := messages[i]
		log.Info("message:", msg.MsgID, im.Command(msg.Cmd))
		m := &im.Message{cmd: int(msg.Cmd), version: im.DEFAULT_VERSION}
		m.FromData(msg.Raw)
		sk.sync_key = msg.MsgID
		if client.isSender(m, msg.DeviceID) {
			m.flag |= im.MESSAGE_FLAG_SELF
		}
		client.EnqueueMessage(m)
	}

	if gh.LastMsgID < last_id && gh.LastMsgID > 0 {
		sk.sync_key = gh.LastMsgID
		log.Warningf("group:%d client last id:%d server last id:%d", group_id, last_id, gh.LastMsgID)
	}
	client.EnqueueMessage(&im.Message{cmd: im.MSG_SYNC_GROUP_END, body: sk})
}

func (client *GroupClient) HandleGroupSyncKey(group_sync_key *im.GroupSyncKey) {
	if client.uid == 0 {
		return
	}

	group_id := group_sync_key.group_id
	last_id := group_sync_key.sync_key

	log.Info("group sync key:", group_sync_key.sync_key, last_id)
	if last_id > 0 {
		s := &storage.SyncGroupHistory{
			AppID:     client.appid,
			Uid:       client.uid,
			GroupID:   group_id,
			LastMsgID: last_id,
		}
		im.group_sync_c <- s
	}
}

func (client *GroupClient) HandleMessage(msg *im.Message) {
	switch msg.cmd {
	case im.MSG_GROUP_IM:
		client.HandleGroupIMMessage(msg)
	case im.MSG_SYNC_GROUP:
		client.HandleGroupSync(msg.body.(*im.GroupSyncKey))
	case im.MSG_GROUP_SYNC_KEY:
		client.HandleGroupSyncKey(msg.body.(*im.GroupSyncKey))
	}
}