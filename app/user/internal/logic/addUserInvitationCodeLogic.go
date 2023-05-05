package logic

import (
	"context"
	"gorm.io/gorm"
	"imservice/app/user/usermodel"
	"imservice/common/xorm"
	"time"

	"imservice/app/user/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserInvitationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserInvitationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserInvitationCodeLogic {
	return &AddUserInvitationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserInvitationCodeLogic) AddUserInvitationCode(in *pb.AddUserInvitationCodeReq) (*pb.AddUserInvitationCodeResp, error) {
	model := &usermodel.InvitationCode{
		Code:             in.UserInvitationCode.Code,
		Remark:           in.UserInvitationCode.Remark,
		Creator:          in.CommonReq.UserId,
		CreatorType:      in.UserInvitationCode.CreatorType,
		IsEnable:         in.UserInvitationCode.IsEnable,
		SuccessUserCount: 0,
		DefaultConvMode:  in.UserInvitationCode.DefaultConvMode,
		CreateTime:       time.Now().UnixMilli(),
	}
	err := xorm.Transaction(l.svcCtx.Mysql(), func(tx *gorm.DB) error {
		err := model.Insert(tx)
		if err != nil {
			l.Errorf("insert err: %v", err)
		}
		return err
	})
	if err != nil {
		l.Errorf("insert err: %v", err)
		return &pb.AddUserInvitationCodeResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.AddUserInvitationCodeResp{}, nil
}
