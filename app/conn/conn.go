package main

import (
	"flag"
	"fmt"
	"imservice/app/mgmt/mgmtservice"

	"imservice/app/conn/internal/config"
	"imservice/app/conn/internal/server"
	"imservice/app/conn/internal/svc"
	"imservice/common/pb"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var mgmtRpcAddress = flag.String("a", "127.0.0.1:6708", "mgmt rpc address")

func main() {
	flag.Parse()

	var c config.Config
	mgmtservice.MustLoadConfig(*mgmtRpcAddress, "conn", &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewConnServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterConnServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	connServer := server.NewConnServer(ctx)
	go connServer.Start()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
