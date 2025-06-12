package mqs

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/order/internal/logic"
	"github.com/ZChen470/eShop/rpc/order/internal/svc"
)

type CheckResult struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckResult(ctx context.Context, svcCtx *svc.ServiceContext) *CheckResult {
	return &CheckResult{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckResult) Consume(ctx context.Context, key, val string) error {
	// 修改订单状态等待用户
	orderId, err := strconv.Atoi(key)
	if err != nil {
		return fmt.Errorf("订单ID转换失败: %v", err)
	}
	if val == "库存验证失败" {
		if err := l.svcCtx.DB.Model(&model.Order{}).
			Where("order_id = ?", orderId).Update("status", logic.OrderStatusUnderstock).Error; err != nil {
			return fmt.Errorf("修改订单状态失败: %v", err)
		}
	} else if val == "库存验证成功" {
		if err := l.svcCtx.DB.Model(&model.Order{}).
			Where("order_id = ?", orderId).Update("status", logic.OrderStatusStockConfirmed).Error; err != nil {
			return fmt.Errorf("修改订单状态失败: %v", err)
		}
	}
	return nil
}
