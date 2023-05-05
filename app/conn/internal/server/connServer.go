package server

import (
	"imservice/app/conn/internal/logic"
	"imservice/app/conn/internal/logic/conngateway"
	"imservice/app/conn/internal/server/route"
	"imservice/app/conn/internal/server/tcp"
	"imservice/app/conn/internal/server/ws"
	"imservice/app/conn/internal/svc"
	"imservice/app/conn/internal/types"
)

type ConnServer struct {
	svcCtx  *svc.ServiceContext
	Servers []types.IServer
}

func NewConnServer(svcCtx *svc.ServiceContext) *ConnServer {
	s := &ConnServer{
		svcCtx: svcCtx,
	}
	logic.InitConnLogic(svcCtx)
	l := logic.GetConnLogic()

	var servers = []types.IServer{ws.NewServer(svcCtx), tcp.NewServer(svcCtx)}
	for _, server := range servers {
		server.SetBeforeConnect(l.BeforeConnect)
		server.SetAddSubscriber(l.AddSubscriber)
		server.SetDeleteSubscriber(l.DeleteSubscriber)
		server.SetOnReceive(l.OnReceive)
	}
	s.Servers = servers
	{
		conngateway.Init(s.svcCtx)
		route.RegisterAppMgmt(s.svcCtx)
		route.RegisterConn(s.svcCtx)
		route.RegisterGroup(s.svcCtx)
		route.RegisterIm(s.svcCtx)
		route.RegisterMsg(s.svcCtx)
		route.RegisterNotice(s.svcCtx)
		route.RegisterRelation(s.svcCtx)
		route.RegisterUser(s.svcCtx)
		conngateway.PrintRoutes()
	}
	go l.Stats()
	return s
}

func (s *ConnServer) Start() {
	for _, server := range s.Servers {
		go server.Start()
	}
}
