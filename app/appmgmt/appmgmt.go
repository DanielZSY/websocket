package main

import (
	"flag"
	"fmt"
	"imservice/app/appmgmt/internal/config"
	"imservice/app/appmgmt/internal/logic"
	"imservice/app/appmgmt/internal/server"
	"imservice/app/appmgmt/internal/svc"
	"imservice/app/mgmt/mgmtservice"
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
	mgmtservice.MustLoadConfig(*mgmtRpcAddress, "appmgmt", &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewAppMgmtServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAppMgmtServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logic.NewStatsLogic(ctx).Start()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
