package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"imservice/app/group/groupservice"
	"imservice/app/im/imservice"
	msgservice "imservice/app/msg/msgService"
	"imservice/app/notice/internal/config"
	"imservice/app/notice/noticemodel"
	"imservice/app/relation/relationservice"
	"imservice/app/user/userservice"
	"imservice/common/xconf"
	"imservice/common/xorm"
)

type ServiceContext struct {
	Config          config.Config
	zedis           *redis.Redis
	mysql           *gorm.DB
	imService       imservice.ImService
	relationService relationservice.RelationService
	groupService    groupservice.GroupService
	userService     userservice.UserService
	msgService      msgservice.MsgService
	ConfigMgr       *xconf.ConfigMgr
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config: c,
	}
	s.Mysql().AutoMigrate(
		noticemodel.Notice{},
		noticemodel.NoticeAckRecord{},
		noticemodel.NoticeMaxConvAutoId{},
		xorm.HashKv{},
	)
	s.ConfigMgr = xconf.NewConfigMgr(s.Mysql(), s.Redis(), "system")
	return s
}

func (s *ServiceContext) Redis() *redis.Redis {
	if s.zedis == nil {
		s.zedis = s.Config.Redis.NewRedis()
	}
	return s.zedis
}

func (s *ServiceContext) Mysql() *gorm.DB {
	if s.mysql == nil {
		s.mysql = xorm.NewClient(s.Config.Mysql)
	}
	return s.mysql
}

func (s *ServiceContext) ImService() imservice.ImService {
	if s.imService == nil {
		s.imService = imservice.NewImService(zrpc.MustNewClient(s.Config.ImRpc))
	}
	return s.imService
}

func (s *ServiceContext) RelationService() relationservice.RelationService {
	if s.relationService == nil {
		s.relationService = relationservice.NewRelationService(zrpc.MustNewClient(s.Config.RelationRpc))
	}
	return s.relationService
}

func (s *ServiceContext) GroupService() groupservice.GroupService {
	if s.groupService == nil {
		s.groupService = groupservice.NewGroupService(zrpc.MustNewClient(s.Config.GroupRpc))
	}
	return s.groupService
}

func (s *ServiceContext) UserService() userservice.UserService {
	if s.userService == nil {
		s.userService = userservice.NewUserService(zrpc.MustNewClient(s.Config.UserRpc))
	}
	return s.userService
}

func (s *ServiceContext) MsgService() msgservice.MsgService {
	if s.msgService == nil {
		s.msgService = msgservice.NewMsgService(zrpc.MustNewClient(s.Config.MsgRpc))
	}
	return s.msgService
}
