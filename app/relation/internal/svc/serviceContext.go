package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"imservice/app/im/imservice"
	msgservice "imservice/app/msg/msgService"
	"imservice/app/notice/noticeservice"
	"imservice/app/relation/internal/config"
	"imservice/app/relation/relationmodel"
	"imservice/app/user/userservice"
	"imservice/common/i18n"
	"imservice/common/xconf"
	"imservice/common/xorm"
)

type ServiceContext struct {
	Config        config.Config
	zedis         *redis.Redis
	mysql         *gorm.DB
	imService     imservice.ImService
	userService   userservice.UserService
	msgService    msgservice.MsgService
	noticeService noticeservice.NoticeService
	ConfigMgr     *xconf.ConfigMgr
	*i18n.I18N
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config: c,
	}
	s.ConfigMgr = xconf.NewConfigMgr(s.Mysql(), s.Redis(), "system")
	s.I18N = i18n.NewI18N(s.Mysql())
	s.Mysql().AutoMigrate(
		relationmodel.Friend{},
		relationmodel.Blacklist{},
		relationmodel.RequestAddFriend{},
	)
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

func (s *ServiceContext) NoticeService() noticeservice.NoticeService {
	if s.noticeService == nil {
		s.noticeService = noticeservice.NewNoticeService(zrpc.MustNewClient(s.Config.NoticeRpc))
	}
	return s.noticeService
}
