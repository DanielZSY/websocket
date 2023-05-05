package logic

import (
	"context"
	"imservice/app/mgmt/mgmtmodel"

	"imservice/app/mgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMSRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMSRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMSRoleLogic {
	return &DeleteMSRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMSRoleLogic) DeleteMSRole(in *pb.DeleteMSRoleReq) (*pb.DeleteMSRoleResp, error) {
	err := l.svcCtx.Mysql().Model(&mgmtmodel.Role{}).Where("id in (?)", in.Ids).Delete(&mgmtmodel.Role{}).Error
	if err != nil {
		l.Errorf("delete error: %v", err)
	}
	return &pb.DeleteMSRoleResp{}, err
}
