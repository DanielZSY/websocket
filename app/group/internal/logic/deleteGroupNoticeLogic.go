package logic

import (
	"context"

	"imservice/app/group/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupNoticeLogic {
	return &DeleteGroupNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteGroupNotice 删除群公告
func (l *DeleteGroupNoticeLogic) DeleteGroupNotice(in *pb.DeleteGroupNoticeReq) (*pb.DeleteGroupNoticeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteGroupNoticeResp{}, nil
}
