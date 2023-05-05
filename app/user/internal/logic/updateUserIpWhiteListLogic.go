package logic

import (
	"context"
	"imservice/app/user/usermodel"

	"imservice/app/user/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserIpWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserIpWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserIpWhiteListLogic {
	return &UpdateUserIpWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserIpWhiteListLogic) UpdateUserIpWhiteList(in *pb.UpdateUserIpWhiteListReq) (*pb.UpdateUserIpWhiteListResp, error) {
	// 查询原模型
	model := &usermodel.IpWhiteList{}
	err := l.svcCtx.Mysql().Model(model).Where("id = ?", in.UserIpList.Id).First(model).Error
	if err != nil {
		l.Errorf("查询失败: %v", err)
		return &pb.UpdateUserIpWhiteListResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	updateMap := map[string]interface{}{}
	{
		updateMap["platform"] = in.UserIpList.Platform
		updateMap["startIp"] = in.UserIpList.StartIp
		updateMap["endIp"] = in.UserIpList.EndIp
		updateMap["remark"] = in.UserIpList.Remark
		updateMap["userId"] = in.UserIpList.UserId
		updateMap["isEnable"] = in.UserIpList.IsEnable
	}
	if len(updateMap) > 0 {
		err = l.svcCtx.Mysql().Model(model).Where("id = ?", in.UserIpList.Id).Updates(updateMap).Error
		if err != nil {
			l.Errorf("更新失败: %v", err)
			return &pb.UpdateUserIpWhiteListResp{CommonResp: pb.NewRetryErrorResp()}, err
		}
	}
	return &pb.UpdateUserIpWhiteListResp{}, nil
}