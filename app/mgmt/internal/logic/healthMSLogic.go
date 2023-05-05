package logic

import (
	"context"

	"imservice/app/mgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthMSLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHealthMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthMSLogic {
	return &HealthMSLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HealthMSLogic) HealthMS(in *pb.CommonReq) (*pb.HealthMSResp, error) {
	// todo: add your logic here and delete this line

	return &pb.HealthMSResp{}, nil
}
