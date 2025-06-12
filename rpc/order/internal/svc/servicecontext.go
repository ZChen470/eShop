package svc

import (
	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/basket/basketclient"
	"github.com/ZChen470/eShop/rpc/order/internal/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	BasketRpcClient basketclient.Basket
	KqPusherClient  *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.DataSource
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logx.Error("数据库连接建立失败", err)
	}
	err = db.AutoMigrate(&model.Order{}, &model.OrderItem{})
	if err != nil {
		logx.Error("数据库迁移失败", err)
	}
	return &ServiceContext{
		Config:          c,
		DB:              db,
		BasketRpcClient: basketclient.NewBasket(zrpc.MustNewClient(c.Basket)),
		KqPusherClient:  kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
