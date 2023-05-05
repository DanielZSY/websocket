package route

import (
	"imservice/app/conn/internal/logic/conngateway"
	"imservice/app/conn/internal/svc"
	"imservice/common/pb"
)

func RegisterIm(svcCtx *svc.ServiceContext) {
	// im
	{
		// UpdateConvSettingReq UpdateConvSettingResp
		{
			route := conngateway.Route[*pb.UpdateConvSettingReq, *pb.UpdateConvSettingResp]{
				NewRequest: func() *pb.UpdateConvSettingReq {
					return &pb.UpdateConvSettingReq{}
				},
				Do: svcCtx.ImService().UpdateConvSetting,
			}
			conngateway.AddRoute("/v1/im/updateConvSetting", route)
		}
		// GetConvSettingReq GetConvSettingResp
		{
			route := conngateway.Route[*pb.GetConvSettingReq, *pb.GetConvSettingResp]{
				NewRequest: func() *pb.GetConvSettingReq {
					return &pb.GetConvSettingReq{}
				},
				Do: svcCtx.ImService().GetConvSetting,
			}
			conngateway.AddRoute("/v1/im/getConvSetting", route)
		}
	}
}
