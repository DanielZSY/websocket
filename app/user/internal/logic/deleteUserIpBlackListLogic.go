package logic

import (
	"context"
	"imservice/app/user/usermodel"

	"imservice/app/user/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserIpBlackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserIpBlackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserIpBlackListLogic {
	return &DeleteUserIpBlackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserIpBlackListLogic) DeleteUserIpBlackList(in *pb.DeleteUserIpBlackListReq) (*pb.DeleteUserIpBlackListResp, error) {
	model := &usermodel.IpBlackList{}
	err := l.svcCtx.Mysql().Model(model).Where("id in (?)", in.Ids).Delete(model).Error
	if err != nil {
		l.Errorf("delete error: %v", err)
		return &pb.DeleteUserIpBlackListResp{
			CommonResp: pb.NewRetryErrorResp(),
		}, err
	}
	return &pb.DeleteUserIpBlackListResp{}, nil
}
