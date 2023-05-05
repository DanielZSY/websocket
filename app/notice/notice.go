package main

import (
	"flag"
	"fmt"
	"imservice/app/mgmt/mgmtservice"

	"imservice/app/notice/internal/config"
	"imservice/app/notice/internal/server"
	"imservice/app/notice/internal/svc"
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
	mgmtservice.MustLoadConfig(*mgmtRpcAddress, "notice", &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewNoticeServiceServer(ctx)
	svr.Start()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterNoticeServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
