package logic

import (
	"context"
	"imservice/app/mgmt/mgmtmodel"
	"time"

	"imservice/app/mgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMSIpWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMSIpWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMSIpWhiteListLogic {
	return &AddMSIpWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMSIpWhiteListLogic) AddMSIpWhiteList(in *pb.AddMSIpWhiteListReq) (*pb.AddMSIpWhiteListResp, error) {
	model := &mgmtmodel.MSIPWhitelist{
		Id:         mgmtmodel.GetId(l.svcCtx.Mysql(), &mgmtmodel.MSIPWhitelist{}, 1000),
		StartIp:    in.IpWhiteList.StartIp,
		EndIp:      in.IpWhiteList.EndIp,
		Remark:     in.IpWhiteList.Remark,
		UserId:     in.IpWhiteList.UserId,
		IsEnable:   in.IpWhiteList.IsEnable,
		CreateTime: time.Now().UnixMilli(),
	}
	err := l.svcCtx.Mysql().Model(model).Create(model).Error
	if err != nil {
		l.Errorf("添加失败: %v", err)
		return &pb.AddMSIpWhiteListResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	return &pb.AddMSIpWhiteListResp{}, nil
}
