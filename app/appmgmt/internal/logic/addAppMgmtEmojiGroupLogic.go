package logic

import (
	"context"
	"gorm.io/gorm"
	"imservice/app/appmgmt/appmgmtmodel"
	"imservice/common/xorm"
	"time"

	"imservice/app/appmgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAppMgmtEmojiGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAppMgmtEmojiGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAppMgmtEmojiGroupLogic {
	return &AddAppMgmtEmojiGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAppMgmtEmojiGroupLogic) AddAppMgmtEmojiGroup(in *pb.AddAppMgmtEmojiGroupReq) (*pb.AddAppMgmtEmojiGroupResp, error) {
	model := &appmgmtmodel.EmojiGroup{
		Name:       in.AppMgmtEmojiGroup.Name,
		CoverId:    "",
		IsEnable:   in.AppMgmtEmojiGroup.IsEnable,
		CreateTime: time.Now().UnixMilli(),
	}
	err := xorm.Transaction(l.svcCtx.Mysql(), func(tx *gorm.DB) error {
		err := model.Insert(tx)
		if err != nil {
			l.Errorf("insert err: %v", err)
		}
		return err
	})
	if err != nil {
		l.Errorf("insert err: %v", err)
		return &pb.AddAppMgmtEmojiGroupResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.AddAppMgmtEmojiGroupResp{}, nil
}
