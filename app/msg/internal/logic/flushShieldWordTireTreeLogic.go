package logic

import (
	"context"

	"imservice/app/msg/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlushShieldWordTireTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlushShieldWordTireTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlushShieldWordTireTreeLogic {
	return &FlushShieldWordTireTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FlushShieldWordTireTree 刷新屏蔽词
func (l *FlushShieldWordTireTreeLogic) FlushShieldWordTireTree(in *pb.FlushShieldWordTireTreeReq) (*pb.FlushShieldWordTireTreeResp, error) {
	ShieldWordTrieTreeInstance.Flush()
	return &pb.FlushShieldWordTireTreeResp{}, nil
}
