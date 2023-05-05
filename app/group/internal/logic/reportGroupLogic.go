package logic

import (
	"context"
	"imservice/app/group/groupmodel"
	"imservice/common/utils"
	"time"

	"imservice/app/group/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReportGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportGroupLogic {
	return &ReportGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReportGroup 举报群
func (l *ReportGroupLogic) ReportGroup(in *pb.ReportGroupReq) (*pb.ReportGroupResp, error) {
	model := &groupmodel.ReportRecord{
		Id:            utils.GenId(),
		ReporterId:    in.CommonReq.UserId,
		ReportedId:    in.GroupId,
		ReportType:    "",
		ReportContent: in.Reason,
		ReportImages:  make([]string, 0),
		ReportTime:    time.Now().UnixMilli(),
		ReportStatus:  "",
		HandleTime:    0,
		HandlerId:     "",
	}
	l.svcCtx.Mysql().Create(model)
	return &pb.ReportGroupResp{}, nil
}
