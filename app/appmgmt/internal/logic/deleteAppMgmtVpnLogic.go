package logic

import (
	"context"
	"imservice/app/appmgmt/appmgmtmodel"

	"imservice/app/appmgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppMgmtVpnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAppMgmtVpnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppMgmtVpnLogic {
	return &DeleteAppMgmtVpnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAppMgmtVpnLogic) DeleteAppMgmtVpn(in *pb.DeleteAppMgmtVpnReq) (*pb.DeleteAppMgmtVpnResp, error) {
	model := &appmgmtmodel.Vpn{}
	err := l.svcCtx.Mysql().Model(model).Where("id in (?)", in.Ids).Delete(model).Error
	if err != nil {
		l.Errorf("delete error: %v", err)
		return &pb.DeleteAppMgmtVpnResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.DeleteAppMgmtVpnResp{}, nil
}