package logic

import (
	"context"
	"imservice/app/appmgmt/appmgmtmodel"

	"imservice/app/appmgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppMgmtEmojiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAppMgmtEmojiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppMgmtEmojiLogic {
	return &DeleteAppMgmtEmojiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAppMgmtEmojiLogic) DeleteAppMgmtEmoji(in *pb.DeleteAppMgmtEmojiReq) (*pb.DeleteAppMgmtEmojiResp, error) {
	model := &appmgmtmodel.Emoji{}
	err := l.svcCtx.Mysql().Model(model).Where("id in (?)", in.Ids).Delete(model).Error
	if err != nil {
		l.Errorf("delete error: %v", err)
		return &pb.DeleteAppMgmtEmojiResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.DeleteAppMgmtEmojiResp{}, nil
}
