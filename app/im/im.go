package main

import (
	"flag"
	"fmt"
	"imservice/app/im/internal/logic"
	"imservice/app/mgmt/mgmtservice"

	"imservice/app/im/internal/config"
	"imservice/app/im/internal/server"
	"imservice/app/im/internal/svc"
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
	mgmtservice.MustLoadConfig(*mgmtRpcAddress, "im", &c)
	ctx := svc.NewServiceContext(c)
	logic.InitAllIpBlackList(ctx)

	svr := server.NewImServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterImServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
