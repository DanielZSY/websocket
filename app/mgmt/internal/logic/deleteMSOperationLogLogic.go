package logic

import (
	"context"
	"imservice/app/mgmt/mgmtmodel"

	"imservice/app/mgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMSOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMSOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMSOperationLogLogic {
	return &DeleteMSOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMSOperationLogLogic) DeleteMSOperationLog(in *pb.DeleteMSOperationLogReq) (*pb.DeleteMSOperationLogResp, error) {
	err := l.svcCtx.Mysql().Model(&mgmtmodel.OperationLog{}).Where("id in (?)", in.Ids).Delete(&mgmtmodel.OperationLog{}).Error
	if err != nil {
		l.Errorf("delete error: %v", err)
	}
	return &pb.DeleteMSOperationLogResp{}, err
}
