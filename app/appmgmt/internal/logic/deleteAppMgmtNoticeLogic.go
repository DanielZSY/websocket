package logic

import (
	"context"
	"imservice/app/appmgmt/appmgmtmodel"

	"imservice/app/appmgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppMgmtNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAppMgmtNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppMgmtNoticeLogic {
	return &DeleteAppMgmtNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAppMgmtNoticeLogic) DeleteAppMgmtNotice(in *pb.DeleteAppMgmtNoticeReq) (*pb.DeleteAppMgmtNoticeResp, error) {
	model := &appmgmtmodel.Notice{}
	err := l.svcCtx.Mysql().Model(model).Where("id in (?)", in.Ids).Delete(model).Error
	if err != nil {
		l.Errorf("delete error: %v", err)
		return &pb.DeleteAppMgmtNoticeResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.DeleteAppMgmtNoticeResp{}, nil
}
