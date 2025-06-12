package svc

import (
	"github.com/ZChen470/eShop/api/internal/config"
	"github.com/ZChen470/eShop/rpc/basket/basketclient"
	"github.com/ZChen470/eShop/rpc/catalog/catalogclient"
	"github.com/ZChen470/eShop/rpc/identify/identifyclient"
	"github.com/ZChen470/eShop/rpc/order/ordering"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CatalogRpcClient  catalogclient.Catalog
	BasketRpcClient   basketclient.Basket
	IdentifyRpcClient identifyclient.Identify
	OrderRpcClient    ordering.Ordering
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,

		CatalogRpcClient:  catalogclient.NewCatalog(zrpc.MustNewClient(c.Catalog)),
		BasketRpcClient:   basketclient.NewBasket(zrpc.MustNewClient(c.Basket)),
		IdentifyRpcClient: identifyclient.NewIdentify(zrpc.MustNewClient(c.Catalog)),
		OrderRpcClient:    ordering.NewOrdering(zrpc.MustNewClient(c.Catalog)),
	}
}
