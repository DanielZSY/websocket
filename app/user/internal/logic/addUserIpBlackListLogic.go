package logic

import (
	"context"
	"gorm.io/gorm"
	"imservice/app/user/internal/svc"
	"imservice/app/user/usermodel"
	"imservice/common/pb"
	"imservice/common/xorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserIpBlackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserIpBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserIpBlackListLogic {
	return &AddUserIpBlackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserIpBlackListLogic) AddUserIpBlackList(in *pb.AddUserIpBlackListReq) (*pb.AddUserIpBlackListResp, error) {
	model := &usermodel.IpBlackList{
		IpList: usermodel.IpList{
			Id:         usermodel.GetId(l.svcCtx.Mysql(), &usermodel.IpBlackList{}, 10000),
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
		return &pb.AddUserIpBlackListResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.AddUserIpBlackListResp{}, nil
}
