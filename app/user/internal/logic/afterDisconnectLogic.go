package logic

import (
	"context"

	"imservice/app/user/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AfterDisconnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAfterDisconnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AfterDisconnectLogic {
	return &AfterDisconnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AfterDisconnect conn hook
func (l *AfterDisconnectLogic) AfterDisconnect(in *pb.AfterDisconnectReq) (*pb.CommonResp, error) {
	// nothing to do
	return &pb.CommonResp{}, nil
}
