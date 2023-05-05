// Code generated by goctl. DO NOT EDIT!
// Source: im.proto

package imservice

import (
	"context"

	"imservice/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchGetUserLatestConnReq  = pb.BatchGetUserLatestConnReq
	BatchGetUserLatestConnResp = pb.BatchGetUserLatestConnResp
	BeforeConnectReq           = pb.BeforeConnectReq
	BeforeConnectResp          = pb.BeforeConnectResp
	BeforeRequestReq           = pb.BeforeRequestReq
	BeforeRequestResp          = pb.BeforeRequestResp
	ConvSetting                = pb.ConvSetting
	GetAllConvIdOfUserReq      = pb.GetAllConvIdOfUserReq
	GetAllConvIdOfUserResp     = pb.GetAllConvIdOfUserResp
	GetConvSettingReq          = pb.GetConvSettingReq
	GetConvSettingResp         = pb.GetConvSettingResp
	GetUserLatestConnReq       = pb.GetUserLatestConnReq
	GetUserLatestConnResp      = pb.GetUserLatestConnResp
	ImMQBody                   = pb.ImMQBody
	MsgNotifyOpt               = pb.MsgNotifyOpt
	UpdateConvSettingReq       = pb.UpdateConvSettingReq
	UpdateConvSettingResp      = pb.UpdateConvSettingResp

	ImService interface {
		BeforeConnect(ctx context.Context, in *BeforeConnectReq, opts ...grpc.CallOption) (*BeforeConnectResp, error)
		AfterConnect(ctx context.Context, in *AfterConnectReq, opts ...grpc.CallOption) (*CommonResp, error)
		AfterDisconnect(ctx context.Context, in *AfterDisconnectReq, opts ...grpc.CallOption) (*CommonResp, error)
		KeepAlive(ctx context.Context, in *KeepAliveReq, opts ...grpc.CallOption) (*KeepAliveResp, error)
		KickUserConn(ctx context.Context, in *KickUserConnReq, opts ...grpc.CallOption) (*KickUserConnResp, error)
		GetUserConn(ctx context.Context, in *GetUserConnReq, opts ...grpc.CallOption) (*GetUserConnResp, error)
		BeforeRequest(ctx context.Context, in *BeforeRequestReq, opts ...grpc.CallOption) (*BeforeRequestResp, error)
		GetUserLatestConn(ctx context.Context, in *GetUserLatestConnReq, opts ...grpc.CallOption) (*GetUserLatestConnResp, error)
		BatchGetUserLatestConn(ctx context.Context, in *BatchGetUserLatestConnReq, opts ...grpc.CallOption) (*BatchGetUserLatestConnResp, error)
		SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
		GetAllConvIdOfUser(ctx context.Context, in *GetAllConvIdOfUserReq, opts ...grpc.CallOption) (*GetAllConvIdOfUserResp, error)
		UpdateConvSetting(ctx context.Context, in *UpdateConvSettingReq, opts ...grpc.CallOption) (*UpdateConvSettingResp, error)
		GetConvSetting(ctx context.Context, in *GetConvSettingReq, opts ...grpc.CallOption) (*GetConvSettingResp, error)
	}

	defaultImService struct {
		cli zrpc.Client
	}
)

func NewImService(cli zrpc.Client) ImService {
	return &defaultImService{
		cli: cli,
	}
}

func (m *defaultImService) BeforeConnect(ctx context.Context, in *BeforeConnectReq, opts ...grpc.CallOption) (*BeforeConnectResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.BeforeConnect(ctx, in, opts...)
}

func (m *defaultImService) AfterConnect(ctx context.Context, in *AfterConnectReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.AfterConnect(ctx, in, opts...)
}

func (m *defaultImService) AfterDisconnect(ctx context.Context, in *AfterDisconnectReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.AfterDisconnect(ctx, in, opts...)
}

func (m *defaultImService) KeepAlive(ctx context.Context, in *KeepAliveReq, opts ...grpc.CallOption) (*KeepAliveResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.KeepAlive(ctx, in, opts...)
}

func (m *defaultImService) KickUserConn(ctx context.Context, in *KickUserConnReq, opts ...grpc.CallOption) (*KickUserConnResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.KickUserConn(ctx, in, opts...)
}

func (m *defaultImService) GetUserConn(ctx context.Context, in *GetUserConnReq, opts ...grpc.CallOption) (*GetUserConnResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.GetUserConn(ctx, in, opts...)
}

func (m *defaultImService) BeforeRequest(ctx context.Context, in *BeforeRequestReq, opts ...grpc.CallOption) (*BeforeRequestResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.BeforeRequest(ctx, in, opts...)
}

func (m *defaultImService) GetUserLatestConn(ctx context.Context, in *GetUserLatestConnReq, opts ...grpc.CallOption) (*GetUserLatestConnResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.GetUserLatestConn(ctx, in, opts...)
}

func (m *defaultImService) BatchGetUserLatestConn(ctx context.Context, in *BatchGetUserLatestConnReq, opts ...grpc.CallOption) (*BatchGetUserLatestConnResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.BatchGetUserLatestConn(ctx, in, opts...)
}

func (m *defaultImService) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.SendMsg(ctx, in, opts...)
}

func (m *defaultImService) GetAllConvIdOfUser(ctx context.Context, in *GetAllConvIdOfUserReq, opts ...grpc.CallOption) (*GetAllConvIdOfUserResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.GetAllConvIdOfUser(ctx, in, opts...)
}

func (m *defaultImService) UpdateConvSetting(ctx context.Context, in *UpdateConvSettingReq, opts ...grpc.CallOption) (*UpdateConvSettingResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.UpdateConvSetting(ctx, in, opts...)
}

func (m *defaultImService) GetConvSetting(ctx context.Context, in *GetConvSettingReq, opts ...grpc.CallOption) (*GetConvSettingResp, error) {
	client := pb.NewImServiceClient(m.cli.Conn())
	return client.GetConvSetting(ctx, in, opts...)
}
