// Code generated by goctl. DO NOT EDIT!
// Source: im.proto

package server

import (
	"context"

	"imservice/app/im/internal/logic"
	"imservice/app/im/internal/svc"
	"imservice/common/pb"
)

type ImServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedImServiceServer
}

func NewImServiceServer(svcCtx *svc.ServiceContext) *ImServiceServer {
	return &ImServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ImServiceServer) BeforeConnect(ctx context.Context, in *pb.BeforeConnectReq) (*pb.BeforeConnectResp, error) {
	l := logic.NewBeforeConnectLogic(ctx, s.svcCtx)
	return l.BeforeConnect(in)
}

func (s *ImServiceServer) AfterConnect(ctx context.Context, in *pb.AfterConnectReq) (*pb.CommonResp, error) {
	l := logic.NewAfterConnectLogic(ctx, s.svcCtx)
	return l.AfterConnect(in)
}

func (s *ImServiceServer) AfterDisconnect(ctx context.Context, in *pb.AfterDisconnectReq) (*pb.CommonResp, error) {
	l := logic.NewAfterDisconnectLogic(ctx, s.svcCtx)
	return l.AfterDisconnect(in)
}

func (s *ImServiceServer) KeepAlive(ctx context.Context, in *pb.KeepAliveReq) (*pb.KeepAliveResp, error) {
	l := logic.NewKeepAliveLogic(ctx, s.svcCtx)
	return l.KeepAlive(in)
}

func (s *ImServiceServer) KickUserConn(ctx context.Context, in *pb.KickUserConnReq) (*pb.KickUserConnResp, error) {
	l := logic.NewKickUserConnLogic(ctx, s.svcCtx)
	return l.KickUserConn(in)
}

func (s *ImServiceServer) GetUserConn(ctx context.Context, in *pb.GetUserConnReq) (*pb.GetUserConnResp, error) {
	l := logic.NewGetUserConnLogic(ctx, s.svcCtx)
	return l.GetUserConn(in)
}

func (s *ImServiceServer) BeforeRequest(ctx context.Context, in *pb.BeforeRequestReq) (*pb.BeforeRequestResp, error) {
	l := logic.NewBeforeRequestLogic(ctx, s.svcCtx)
	return l.BeforeRequest(in)
}

func (s *ImServiceServer) GetUserLatestConn(ctx context.Context, in *pb.GetUserLatestConnReq) (*pb.GetUserLatestConnResp, error) {
	l := logic.NewGetUserLatestConnLogic(ctx, s.svcCtx)
	return l.GetUserLatestConn(in)
}

func (s *ImServiceServer) BatchGetUserLatestConn(ctx context.Context, in *pb.BatchGetUserLatestConnReq) (*pb.BatchGetUserLatestConnResp, error) {
	l := logic.NewBatchGetUserLatestConnLogic(ctx, s.svcCtx)
	return l.BatchGetUserLatestConn(in)
}

func (s *ImServiceServer) SendMsg(ctx context.Context, in *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	l := logic.NewSendMsgLogic(ctx, s.svcCtx)
	return l.SendMsg(in)
}

func (s *ImServiceServer) GetAllConvIdOfUser(ctx context.Context, in *pb.GetAllConvIdOfUserReq) (*pb.GetAllConvIdOfUserResp, error) {
	l := logic.NewGetAllConvIdOfUserLogic(ctx, s.svcCtx)
	return l.GetAllConvIdOfUser(in)
}

func (s *ImServiceServer) UpdateConvSetting(ctx context.Context, in *pb.UpdateConvSettingReq) (*pb.UpdateConvSettingResp, error) {
	l := logic.NewUpdateConvSettingLogic(ctx, s.svcCtx)
	return l.UpdateConvSetting(in)
}

func (s *ImServiceServer) GetConvSetting(ctx context.Context, in *pb.GetConvSettingReq) (*pb.GetConvSettingResp, error) {
	l := logic.NewGetConvSettingLogic(ctx, s.svcCtx)
	return l.GetConvSetting(in)
}