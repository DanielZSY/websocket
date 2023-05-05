package logic

import (
	"context"
	"imservice/app/appmgmt/appmgmtmodel"

	"imservice/app/appmgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppMgmtVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAppMgmtVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppMgmtVersionLogic {
	return &DeleteAppMgmtVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAppMgmtVersionLogic) DeleteAppMgmtVersion(in *pb.DeleteAppMgmtVersionReq) (*pb.DeleteAppMgmtVersionResp, error) {
	model := &appmgmtmodel.Version{}
	err := l.svcCtx.Mysql().Model(model).Where("id in (?)", in.Ids).Delete(model).Error
	if err != nil {
		l.Errorf("delete error: %v", err)
		return &pb.DeleteAppMgmtVersionResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.DeleteAppMgmtVersionResp{}, nil
}
