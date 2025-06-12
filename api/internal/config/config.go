package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// PGSqlConfig
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	// Rpc Client Config
	Catalog  zrpc.RpcClientConf
	Ordering zrpc.RpcClientConf
	Basket   zrpc.RpcClientConf
	Identity zrpc.RpcClientConf
}

// type PGSqlConfig struct {
// 	DataSource string
// }
