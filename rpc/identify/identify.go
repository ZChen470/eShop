package main

import (
	"flag"
	"fmt"

	"github.com/ZChen470/eShop/rpc/identify/identify"
	"github.com/ZChen470/eShop/rpc/identify/internal/config"
	"github.com/ZChen470/eShop/rpc/identify/internal/server"
	"github.com/ZChen470/eShop/rpc/identify/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/identify.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		identify.RegisterIdentifyServer(grpcServer, server.NewIdentifyServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
