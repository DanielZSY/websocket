package logic

import (
	"context"
	"imservice/app/mgmt/mgmtmodel"
	"imservice/common/xorm"
	"time"

	"imservice/app/mgmt/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMSMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMSMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMSMenuLogic {
	return &AddMSMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMSMenuLogic) AddMSMenu(in *pb.AddMSMenuReq) (*pb.AddMSMenuResp, error) {
	model := &mgmtmodel.Menu{
		Id:               mgmtmodel.GetId(l.svcCtx.Mysql(), &mgmtmodel.Menu{}, 10000),
		Pid:              in.Menu.Pid,
		MenuType:         in.Menu.MenuType,
		MenuName:         in.Menu.MenuName,
		MenuIcon:         in.Menu.MenuIcon,
		MenuIconElement2: in.Menu.MenuIconElement2,
		MenuSort:         in.Menu.MenuSort,
		Perms:            in.Menu.Perms,
		Paths:            in.Menu.Paths,
		Component:        in.Menu.Component,
		Selected:         in.Menu.Selected,
		Params:           in.Menu.Params,
		IsCache:          in.Menu.IsCache,
		IsShow:           in.Menu.IsShow,
		IsDisable:        in.Menu.IsDisable,
		CreateTime:       time.Now().UnixMilli(),
		UpdateTime:       time.Now().UnixMilli(),
	}
	err := xorm.InsertOne(l.svcCtx.Mysql(), model)
	if err != nil {
		l.Errorf("AddMSMenu error: %v", err)
		return &pb.AddMSMenuResp{CommonResp: pb.NewRetryErrorResp()}, err
	}
	return &pb.AddMSMenuResp{}, nil
}
