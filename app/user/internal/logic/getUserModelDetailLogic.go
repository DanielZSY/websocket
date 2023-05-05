package logic

import (
	"context"
	"imservice/app/user/usermodel"

	"imservice/app/user/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserModelDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserModelDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserModelDetailLogic {
	return &GetUserModelDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserModelDetailLogic) GetUserModelDetail(in *pb.GetUserModelDetailReq) (*pb.GetUserModelDetailResp, error) {
	// 查询原模型
	model := &usermodel.User{}
	err := l.svcCtx.Mysql().Model(model).Where("id = ?", in.Id).First(model).Error
	if err != nil {
		l.Errorf("查询失败: %v", err)
		return &pb.GetUserModelDetailResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	return &pb.GetUserModelDetailResp{UserModel: model.ToPB()}, nil
}
