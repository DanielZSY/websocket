package logic

import (
	"context"
	"imservice/app/appmgmt/appmgmtmodel"
	"time"

	"imservice/app/appmgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAppMgmtVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAppMgmtVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAppMgmtVersionLogic {
	return &AddAppMgmtVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAppMgmtVersionLogic) AddAppMgmtVersion(in *pb.AddAppMgmtVersionReq) (*pb.AddAppMgmtVersionResp, error) {
	model := &appmgmtmodel.Version{
		Id:          appmgmtmodel.GetId(l.svcCtx.Mysql(), &appmgmtmodel.Version{}, 10000),
		Version:     in.AppMgmtVersion.Version,
		Platform:    in.AppMgmtVersion.Platform,
		Type:        int8(in.AppMgmtVersion.Type),
		Content:     in.AppMgmtVersion.Content,
		DownloadUrl: in.AppMgmtVersion.DownloadUrl,
		CreateTime:  time.Now().UnixMilli(),
	}
	err := model.Insert(l.svcCtx.Mysql())
	if err != nil {
		l.Errorf("insert err: %v", err)
		return &pb.AddAppMgmtVersionResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.AddAppMgmtVersionResp{}, nil
}
