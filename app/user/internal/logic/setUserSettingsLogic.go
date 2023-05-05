package logic

import (
	"context"
	"imservice/app/user/internal/svc"
	"imservice/app/user/usermodel"
	"imservice/common/pb"
	"imservice/common/xorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserSettingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserSettingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserSettingsLogic {
	return &SetUserSettingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserSettingsLogic) SetUserSettings(in *pb.SetUserSettingsReq) (*pb.SetUserSettingsResp, error) {
	for _, setting := range in.Settings {
		model := &usermodel.UserSetting{
			UserId: in.CommonReq.UserId,
			Key:    setting.Key,
			Value:  setting.Value,
		}
		// upsert
		err := xorm.Upsert(l.svcCtx.Mysql(), model, []string{"value"}, []string{"userId", "key"})
		if err != nil {
			l.Errorf("set user setting failed, err: %v", err)
			return &pb.SetUserSettingsResp{CommonResp: pb.NewRetryErrorResp()}, nil
		}
	}
	return &pb.SetUserSettingsResp{}, nil
}
