package logic

import (
	"context"
	"imservice/app/appmgmt/appmgmtmodel"
	"imservice/app/appmgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppMgmtEmojiDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppMgmtEmojiDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppMgmtEmojiDetailLogic {
	return &GetAppMgmtEmojiDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppMgmtEmojiDetailLogic) GetAppMgmtEmojiDetail(in *pb.GetAppMgmtEmojiDetailReq) (*pb.GetAppMgmtEmojiDetailResp, error) {
	// 查询原模型
	model := &appmgmtmodel.Emoji{}
	err := l.svcCtx.Mysql().Model(model).Where("id = ?", in.Id).First(model).Error
	if err != nil {
		l.Errorf("查询失败: %v", err)
		return &pb.GetAppMgmtEmojiDetailResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	return &pb.GetAppMgmtEmojiDetailResp{AppMgmtEmoji: model.ToPB()}, nil
}
