package route

import (
	"imservice/app/conn/internal/logic/conngateway"
	"imservice/app/conn/internal/svc"
	"imservice/common/pb"
)

func RegisterNotice(svcCtx *svc.ServiceContext) {
	// AckNoticeDataReq AckNoticeDataResp
	{
		route := conngateway.Route[*pb.AckNoticeDataReq, *pb.AckNoticeDataResp]{
			NewRequest: func() *pb.AckNoticeDataReq {
				return &pb.AckNoticeDataReq{}
			},
			Do: svcCtx.NoticeService().AckNoticeData,
		}
		conngateway.AddRoute("/v1/notice/ackNoticeData", route)
	}
}
