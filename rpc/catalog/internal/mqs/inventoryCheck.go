package mqs

import (
	"context"
	"encoding/json"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/catalog/internal/svc"
)

type InventoryCheck struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInventoryCheck(ctx context.Context, svcCtx *svc.ServiceContext) *InventoryCheck {
	return &InventoryCheck{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InventoryCheck) Consume(ctx context.Context, key, val string) error {
	// key => orderId
	type product struct {
		ProductId uint  `json:"product_id"`
		Quantity  int32 `json:"quantity"`
	}
	var products []product
	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return err
	}
	for _, v := range products {
		var q int32
		if err := l.svcCtx.DB.Model(&model.Product{}).
			Select("quantity").
			Where("product_id = ?", v.ProductId).
			Find(&q).Error; err != nil {
			return err
		}
		if q < v.Quantity {
			// 发送
			l.svcCtx.KqPusherClient.KPush(l.ctx, key, "库存验证失败")
		}
	}
	l.svcCtx.KqPusherClient.KPush(l.ctx, key, "库存验证成功")
	return nil
}
