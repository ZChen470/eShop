package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Basket zrpc.RpcClientConf
	PGSqlConfig
	KqPusherConf
	KqConsumerConf kq.KqConf
}

type PGSqlConfig struct {
	DataSource string
}

type KqPusherConf struct {
	Brokers []string
	Topic   string
}
