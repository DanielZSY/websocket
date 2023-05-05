package logic

import (
	"context"
	"imservice/app/group/groupmodel"
	"imservice/app/group/internal/svc"
	"imservice/common/pb"
	"imservice/common/xtrace"
)

func getAllGroupManager(ctx context.Context, svcCtx *svc.ServiceContext, groupId string, withOwner bool) ([]*groupmodel.GroupMember, error) {
	var managers []*groupmodel.GroupMember
	var err error
	xtrace.StartFuncSpan(ctx, "getAllGroupManager", func(ctx context.Context) {
		if !withOwner {
			err = svcCtx.Mysql().Model(&groupmodel.GroupMember{}).Where("groupId = ? and role = ?", groupId, pb.GroupRole_MANAGER).Find(&managers).Error
		} else {
			err = svcCtx.Mysql().Model(&groupmodel.GroupMember{}).Where("groupId = ? and (role = ? or role = ?)", groupId, pb.GroupRole_MANAGER, pb.GroupRole_OWNER).Find(&managers).Error
		}
	})
	return managers, err
}
