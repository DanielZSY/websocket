package logic

import (
	"context"
	"imservice/app/mgmt/mgmtmodel"
	"imservice/common/utils"
	"imservice/common/utils/ip2region"
	"imservice/common/xjwt"
	"imservice/common/xpwd"
	"time"

	"imservice/app/mgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginMSLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginMSLogic {
	return &LoginMSLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginMSLogic) LoginMS(in *pb.LoginMSReq) (*pb.LoginMSResp, error) {
	// 查询原模型
	user := &mgmtmodel.User{}
	err := l.svcCtx.Mysql().Model(user).Where("id = ?", in.Id).First(user).Error
	if err != nil {
		l.Errorf("查询用户失败: %v", err)
		return &pb.LoginMSResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	// 用户存在 判断密码是否正确
	if !xpwd.VerifyPwd(in.Password, user.Password, user.PasswordSalt) {
		return &pb.LoginMSResp{CommonResp: pb.NewAlertErrorResp("登录失败", "密码错误")}, nil
	}
	tokenObj := xjwt.GenerateToken(user.Id, "",
		xjwt.WithPlatform(in.CommonReq.Platform),
		xjwt.WithDeviceId(in.CommonReq.DeviceId),
		xjwt.WithDeviceModel(in.CommonReq.DeviceModel),
	)
	err = xjwt.SaveTokenAdmin(l.ctx, l.svcCtx.Redis(), tokenObj)
	if err != nil {
		l.Errorf("save token failed, err: %v", err)
		return &pb.LoginMSResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	region := ip2region.Ip2Region(in.CommonReq.Ip)
	loginRecord := &mgmtmodel.LoginRecord{
		Id:     utils.GenId(),
		UserId: user.Id,
		LoginRecordInfo: mgmtmodel.LoginRecordInfo{
			Time:       time.Now().UnixMilli(),
			Ip:         in.CommonReq.Ip,
			IpCountry:  region.Country,
			IpProvince: region.Province,
			IpCity:     region.City,
			IpISP:      region.ISP,
			UserAgent:  in.CommonReq.UserAgent,
		},
	}
	err = l.svcCtx.Mysql().Model(loginRecord).Create(loginRecord).Error
	if err != nil {
		l.Errorf("保存登录记录失败: %v", err)
	}
	return &pb.LoginMSResp{
		Id:    user.Id,
		Token: tokenObj.Token,
	}, nil
}
