package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	PGSqlConfig

	Auth struct {
		AccessSecret string // 密钥
		AccessExpire int64  // 过期时间 单位秒
	}
}

type PGSqlConfig struct {
	DataSource string
}
