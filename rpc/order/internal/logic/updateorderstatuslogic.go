package logic

import (
	"context"
	"errors"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/order/internal/svc"
	"github.com/ZChen470/eShop/rpc/order/order"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *order.UpdateOrderStatusReq) (*order.CommonResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		var o model.Order
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		}). // 添加 FOR UPDATE 锁 不阻塞等待
			Where("order_id = @order", map[string]interface{}{"order": in.OrderId}).
			Take(&o).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return status.Error(404, "订单不存在")
			}
			return status.Error(500, "查询订单失败")
		}
		switch in.Status {
		case OrderStatusCancelled:
			if o.Status != OrderStatusPending && o.Status != OrderStatusPaid {
				return status.Error(400, "订单状态不支持取消")
			}
		case OrderStatusPaid:
			if o.Status != OrderStatusPending {
				return status.Error(400, "订单状态不支持支付")
			}
		case OrderStatusShipped:
			if o.Status != OrderStatusPaid {
				return status.Error(400, "订单状态不支持发货")
			}
		}

		if err := tx.Model(&o).Update("status", in.Status).Error; err != nil {
			return status.Error(500, "订单状态修改失败")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &order.CommonResp{
		Msg:  "订单状态修改成功",
		Code: 10502,
	}, nil
}
