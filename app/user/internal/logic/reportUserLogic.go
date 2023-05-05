package logic

import (
	"context"
	"imservice/app/user/usermodel"
	"imservice/common/utils"
	"time"

	"imservice/app/user/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReportUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportUserLogic {
	return &ReportUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReportUserLogic) ReportUser(in *pb.ReportUserReq) (*pb.ReportUserResp, error) {
	model := &usermodel.ReportRecord{
		Id:            utils.GenId(),
		ReporterId:    in.CommonReq.UserId,
		ReportedId:    in.UserId,
		ReportType:    "",
		ReportContent: in.Reason,
		ReportImages:  make([]string, 0),
		ReportTime:    time.Now().UnixMilli(),
		ReportStatus:  "",
		HandleTime:    0,
		HandlerId:     "",
	}
	l.svcCtx.Mysql().Create(model)
	return &pb.ReportUserResp{}, nil
}
