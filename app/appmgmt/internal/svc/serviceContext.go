package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"imservice/app/appmgmt/appmgmtmodel"
	"imservice/app/appmgmt/internal/config"
	"imservice/app/group/groupservice"
	"imservice/app/im/imservice"
	"imservice/app/mgmt/mgmtmodel"
	msgservice "imservice/app/msg/msgService"
	"imservice/app/notice/noticeservice"
	"imservice/app/relation/relationservice"
	"imservice/app/user/userservice"
	"imservice/common/utils/ip2region"
	"imservice/common/xconf"
	"imservice/common/xorm"
)

type ServiceContext struct {
	Config          config.Config
	redis           *redis.Redis
	mysql           *gorm.DB
	imService       imservice.ImService
	msgService      msgservice.MsgService
	noticeService   noticeservice.NoticeService
	relationService relationservice.RelationService
	userService     userservice.UserService
	groupService    groupservice.GroupService
	ConfigMgr       *xconf.ConfigMgr
	MsgPodsMgr      *msgservice.MsgPodsMgr
}

func NewServiceContext(c config.Config) *ServiceContext {
	ip2region.Init(c.Ip2RegionUrl)
	s := &ServiceContext{
		Config: c,
	}
	s.MsgPodsMgr = msgservice.NewMsgPodsMgr(c.MsgRpcPod)
	s.ConfigMgr = xconf.NewConfigMgr(s.Mysql(), s.Redis(), "system")
	return s
}

func (s *ServiceContext) Mysql() *gorm.DB {
	if s.mysql == nil {
		s.mysql = xorm.NewClient(s.Config.Mysql)
		s.mysql.AutoMigrate(&appmgmtmodel.AutoIncrement{})
		s.mysql.AutoMigrate(&appmgmtmodel.Config{})
		s.mysql.AutoMigrate(&appmgmtmodel.Emoji{})
		s.mysql.AutoMigrate(&appmgmtmodel.EmojiGroup{})
		s.mysql.AutoMigrate(&appmgmtmodel.Notice{})
		s.mysql.AutoMigrate(&appmgmtmodel.ShieldWord{})
		s.mysql.AutoMigrate(&appmgmtmodel.Version{})
		s.mysql.AutoMigrate(&appmgmtmodel.Vpn{})
		s.mysql.AutoMigrate(&appmgmtmodel.Link{})
		s.mysql.AutoMigrate(&appmgmtmodel.Stats{})
		mgmtmodel.InitData(s.mysql)
	}
	return s.mysql
}

func (s *ServiceContext) ImService() imservice.ImService {
	if s.imService == nil {
		s.imService = imservice.NewImService(zrpc.MustNewClient(s.Config.ImRpc))
	}
	return s.imService
}

func (s *ServiceContext) MsgService() msgservice.MsgService {
	if s.msgService == nil {
		s.msgService = msgservice.NewMsgService(zrpc.MustNewClient(s.Config.MsgRpc))
	}
	return s.msgService
}

func (s *ServiceContext) NoticeService() noticeservice.NoticeService {
	if s.noticeService == nil {
		s.noticeService = noticeservice.NewNoticeService(zrpc.MustNewClient(s.Config.NoticeRpc))
	}
	return s.noticeService
}

func (s *ServiceContext) RelationService() relationservice.RelationService {
	if s.relationService == nil {
		s.relationService = relationservice.NewRelationService(zrpc.MustNewClient(s.Config.RelationRpc))
	}
	return s.relationService
}

func (s *ServiceContext) UserService() userservice.UserService {
	if s.userService == nil {
		s.userService = userservice.NewUserService(zrpc.MustNewClient(s.Config.UserRpc))
	}
	return s.userService
}

func (s *ServiceContext) GroupService() groupservice.GroupService {
	if s.groupService == nil {
		s.groupService = groupservice.NewGroupService(zrpc.MustNewClient(s.Config.GroupRpc))
	}
	return s.groupService
}

func (s *ServiceContext) Redis() *redis.Redis {
	if s.redis == nil {
		s.redis = s.Config.Redis.NewRedis()
	}
	return s.redis
}
