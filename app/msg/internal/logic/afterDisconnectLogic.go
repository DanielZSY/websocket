package logic

import (
	"context"

	"imservice/app/msg/internal/svc"
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

func (l *AfterDisconnectLogic) AfterDisconnect(in *pb.AfterDisconnectReq) (*pb.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CommonResp{}, nil
}
