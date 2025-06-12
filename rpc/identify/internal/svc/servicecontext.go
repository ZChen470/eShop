package svc

import (
	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/identify/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.DataSource
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logx.Error("数据库连接建立失败", err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		logx.Error("数据库迁移失败", err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
