package logic

import (
	"context"
	"imservice/app/user/usermodel"

	"imservice/app/user/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserIpBlackListDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserIpBlackListDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserIpBlackListDetailLogic {
	return &GetUserIpBlackListDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserIpBlackListDetailLogic) GetUserIpBlackListDetail(in *pb.GetUserIpBlackListDetailReq) (*pb.GetUserIpBlackListDetailResp, error) {
	// 查询原模型
	model := &usermodel.IpBlackList{}
	err := l.svcCtx.Mysql().Model(model).Where("id = ?", in.Id).First(model).Error
	if err != nil {
		l.Errorf("查询失败: %v", err)
		return &pb.GetUserIpBlackListDetailResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	return &pb.GetUserIpBlackListDetailResp{UserIpList: model.ToPB()}, nil
}
