package logic

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/order/internal/svc"
	"github.com/ZChen470/eShop/rpc/order/order"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckOutOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckOutOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckOutOrderLogic {
	return &CheckOutOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckOutOrderLogic) CheckOutOrder(in *order.CheckOutOrderReq) (*order.CommonResp, error) {
	//
	userIdVal := l.ctx.Value("userId")
	if userIdVal == nil {
		return nil, status.Error(500, "用户ID无效")
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%v", userIdVal), 10, 64)
	if err != nil {
		return nil, status.Error(500, "用户ID无效")
	}
	// 查询订单商品ID以及数量
	o := new(model.Order)
	if err := l.svcCtx.DB.Preload("Items"). // 不关联查询 Product
						Where("order_id = ? AND user_id = ?", in.OrderId, userId).
						Find(o).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(404, "订单不存在")
		}
		return nil, status.Error(500, "查询订单失败")
	}
	if o.Status != OrderStatusStockConfirmed {
		return nil, status.Error(400, "订单状态错误")
	}

	// 开启事务，扣减订单项对应的商品库存
	tx := l.svcCtx.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 锁住商品并检查库存
	for _, item := range o.Items {
		var product model.Product
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("product_id = ?", item.ProductId).
			First(&product).Error; err != nil {
			tx.Rollback()
			return nil, errors.New("商品查询失败")
		}

		if product.Stock < item.Quantity {
			// 库存不足
			tx.Rollback()
			return &order.CommonResp{
				Msg:  "库存不足",
				Code: 10513,
			}, nil
		}

		// 扣减库存
		if err := tx.Model(&product).
			Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return nil, errors.New("库存扣减失败")
		}
	}

	// 所有商品库存都足够，提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	// 更新订单状态为已支付
	if err := l.svcCtx.DB.Model(o).Update("status", OrderStatusPaid).Error; err != nil {
		return nil, err
	}
	return &order.CommonResp{
		Msg:  "订单支付成功",
		Code: 10503,
	}, nil
}
