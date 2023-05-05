package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"imservice/app/group/groupservice"
	"imservice/app/im/imservice"
	"imservice/app/notice/noticeservice"
	"imservice/app/relation/relationservice"
	"imservice/app/user/internal/config"
	"imservice/app/user/usermodel"
	"imservice/common/i18n"
	"imservice/common/utils/ip2region"
	"imservice/common/utils/xsms"
	"imservice/common/xconf"
	"imservice/common/xorm"
)

type ServiceContext struct {
	Config          config.Config
	zedis           *redis.Redis
	mysql           *gorm.DB
	imService       imservice.ImService
	noticeService   noticeservice.NoticeService
	relationService relationservice.RelationService
	groupService    groupservice.GroupService
	ConfigMgr       *xconf.ConfigMgr
	*i18n.I18N
}

func NewServiceContext(c config.Config) *ServiceContext {
	ip2region.Init(c.Ip2RegionUrl)
	s := &ServiceContext{
		Config: c,
	}
	s.ConfigMgr = xconf.NewConfigMgr(s.Mysql(), s.Redis(), "system")
	s.I18N = i18n.NewI18N(s.Mysql())
	usermodel.InitUserSetting(s.Mysql())
	s.Mysql().AutoMigrate(
		&usermodel.User{},
		&usermodel.UserSetting{},
		&usermodel.UserTmp{},
		&usermodel.LoginRecord{},
		&usermodel.AutoIncrement{},
		&usermodel.DefaultConv{},
		&usermodel.InvitationCode{},
		&usermodel.IpWhiteList{},
		&usermodel.IpBlackList{},
		&usermodel.UserRecycleBin{},
		&usermodel.StatusRecord{},
		&usermodel.ReportRecord{},
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

func (s *ServiceContext) SmsSender() (xsms.SmsSender, error) {
	return xsms.NewSmsSender(s.Config.Sms.ToPb())
}
