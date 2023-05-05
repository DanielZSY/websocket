package logic

import (
	"context"
	"imservice/app/user/usermodel"

	"imservice/app/user/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserDefaultConvLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserDefaultConvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserDefaultConvLogic {
	return &DeleteUserDefaultConvLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserDefaultConvLogic) DeleteUserDefaultConv(in *pb.DeleteUserDefaultConvReq) (*pb.DeleteUserDefaultConvResp, error) {
	model := &usermodel.DefaultConv{}
	err := l.svcCtx.Mysql().Model(model).Where("id in (?)", in.Ids).Delete(model).Error
	if err != nil {
		l.Errorf("delete error: %v", err)
		return &pb.DeleteUserDefaultConvResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.DeleteUserDefaultConvResp{}, nil
}
