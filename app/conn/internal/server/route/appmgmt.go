package route

import (
	"imservice/app/conn/internal/logic/conngateway"
	"imservice/app/conn/internal/svc"
	"imservice/common/pb"
)

func RegisterAppMgmt(svcCtx *svc.ServiceContext) {
	// AppGetAllConfigReq AppGetAllConfigResp
	{
		route := conngateway.Route[*pb.AppGetAllConfigReq, *pb.AppGetAllConfigResp]{
			NewRequest: func() *pb.AppGetAllConfigReq {
				return &pb.AppGetAllConfigReq{}
			},
			Do: svcCtx.AppMgmtService().AppGetAllConfig,
		}
		conngateway.AddRoute("/v1/appmgmt/white/appGetAllConfig", route)
	}
	// GetLatestVersionReq GetLatestVersionResp
	{
		route := conngateway.Route[*pb.GetLatestVersionReq, *pb.GetLatestVersionResp]{
			NewRequest: func() *pb.GetLatestVersionReq {
				return &pb.GetLatestVersionReq{}
			},
			Do: svcCtx.AppMgmtService().GetLatestVersion,
		}
		conngateway.AddRoute("/v1/appmgmt/white/getLatestVersion", route)
	}
}
