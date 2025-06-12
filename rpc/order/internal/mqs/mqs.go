package mqs

import (
	"context"

	"github.com/ZChen470/eShop/rpc/order/internal/config"
	"github.com/ZChen470/eShop/rpc/order/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {

	return []service.Service{
		// Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConsumerConf, NewCheckResult(ctx, svcCtx)),
	}
}
