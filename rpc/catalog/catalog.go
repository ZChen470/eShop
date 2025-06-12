package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ZChen470/eShop/rpc/catalog/catalog"
	"github.com/ZChen470/eShop/rpc/catalog/internal/config"
	"github.com/ZChen470/eShop/rpc/catalog/internal/mqs"
	"github.com/ZChen470/eShop/rpc/catalog/internal/server"
	"github.com/ZChen470/eShop/rpc/catalog/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/catalog.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		catalog.RegisterCatalogServer(grpcServer, server.NewCatalogServer(svcCtx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
