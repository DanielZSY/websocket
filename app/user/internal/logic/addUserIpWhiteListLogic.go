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

type AddUserIpWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserIpWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserIpWhiteListLogic {
	return &AddUserIpWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserIpWhiteListLogic) AddUserIpWhiteList(in *pb.AddUserIpWhiteListReq) (*pb.AddUserIpWhiteListResp, error) {
	model := &usermodel.IpWhiteList{
		IpList: usermodel.IpList{
			Id:         usermodel.GetId(l.svcCtx.Mysql(), &usermodel.IpWhiteList{}, 10000),
			Platform:   in.UserIpList.Platform,
			StartIp:    in.UserIpList.StartIp,
			EndIp:      in.UserIpList.EndIp,
			Remark:     in.UserIpList.Remark,
			UserId:     in.UserIpList.UserId,
			IsEnable:   in.UserIpList.IsEnable,
			CreateTime: time.Now().UnixMilli(),
		},
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
		return &pb.AddUserIpWhiteListResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.AddUserIpWhiteListResp{}, nil
}
