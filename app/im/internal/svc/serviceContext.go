package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"imservice/app/conn/connservice"
	"imservice/app/group/groupservice"
	"imservice/app/im/immodel"
	"imservice/app/im/internal/config"
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
	ConnPodsMgr     *connservice.ConnPodsMgr
	zedis           *redis.Redis
	mysql           *gorm.DB
	msgService      msgservice.MsgService
	relationService relationservice.RelationService
	groupService    groupservice.GroupService
	userService     userservice.UserService
	noticeService   noticeservice.NoticeService
	ConfigMgr       *xconf.ConfigMgr
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
		s.mysql.AutoMigrate(&immodel.ConvSetting{})
	}
	return s.mysql
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

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config: c,
	}
	ip2region.Init(c.Ip2RegionUrl)
	s.ConfigMgr = xconf.NewConfigMgr(s.Mysql(), s.Redis(), "system")
	s.ConnPodsMgr = connservice.NewConnPodsMgr(c.ConnRpc)
	s.Mysql().AutoMigrate(&immodel.UserConnectRecord{})
	return s
}
