package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"imservice/app/appmgmt/appmgmtservice"
	"imservice/app/conn/internal/config"
	"imservice/app/group/groupservice"
	"imservice/app/im/imservice"
	msgservice "imservice/app/msg/msgService"
	"imservice/app/notice/noticeservice"
	"imservice/app/relation/relationservice"
	"imservice/app/user/userservice"
	"imservice/common/utils"
	"imservice/common/utils/ip2region"
)

type ServiceContext struct {
	Config          config.Config
	imService       imservice.ImService
	msgService      msgservice.MsgService
	noticeService   noticeservice.NoticeService
	relationService relationservice.RelationService
	appMgmtService  appmgmtservice.AppMgmtService
	userService     userservice.UserService
	groupService    groupservice.GroupService
	zedis           *redis.Redis
	PodIp           string
}

func NewServiceContext(c config.Config) *ServiceContext {
	ip2region.Init(c.Ip2RegionUrl)
	s := &ServiceContext{
		Config: c,
		PodIp:  utils.GetPodIp(),
	}
	return s
}

func (s *ServiceContext) ImService() imservice.ImService {
	if s.imService == nil {
		s.imService = imservice.NewImService(zrpc.MustNewClient(s.Config.ImRpc))
	}
	return s.imService
}

func (s *ServiceContext) Redis() *redis.Redis {
	if s.zedis == nil {
		s.zedis = s.Config.Redis.NewRedis()
	}
	return s.zedis
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

func (s *ServiceContext) AppMgmtService() appmgmtservice.AppMgmtService {
	if s.appMgmtService == nil {
		s.appMgmtService = appmgmtservice.NewAppMgmtService(zrpc.MustNewClient(s.Config.AppMgmtRpc,
			utils.Zrpc.Options()...))
	}
	return s.appMgmtService
}
