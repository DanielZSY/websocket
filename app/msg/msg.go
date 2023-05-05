package main

import (
	"flag"
	"fmt"
	"imservice/app/mgmt/mgmtservice"
	"imservice/app/msg/internal/logic"

	"imservice/app/msg/internal/config"
	"imservice/app/msg/internal/server"
	"imservice/app/msg/internal/svc"
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
	mgmtservice.MustLoadConfig(*mgmtRpcAddress, "msg", &c)
	ctx := svc.NewServiceContext(c)
	logic.InitShieldWordTrieTree(ctx)

	svr := server.NewMsgServiceServer(ctx)
	svr.Start()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMsgServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
