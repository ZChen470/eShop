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

type CancelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelOrderLogic) CancelOrder(in *order.CancelOrderReq) (*order.CommonResp, error) {
	// todo: add your logic here and delete this line
	userIdVal := l.ctx.Value("userId")
	if userIdVal == nil {
		return nil, fmt.Errorf("用户ID无效")
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%v", userIdVal), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("用户ID无效：%v", err)
	}

	// 查询+修改 原子完成
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		var o model.Order
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		}). // 添加 FOR UPDATE 锁 不阻塞等待
			Where("order_id = @order AND user_id = @user", map[string]interface{}{"order": in.OrderId, "user": userId}).
			Take(&o).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return status.Error(404, "订单不存在")
			}
			return status.Error(500, "查询订单失败")
		}

		if o.Status != OrderStatusPending && o.Status != OrderStatusPaid {
			return status.Error(400, "订单状态不支持取消")
		}
		if err := tx.Model(&o).Update("status", OrderStatusCancelled).Error; err != nil {
			return status.Error(500, "订单取消失败")
		}
		return nil
	})
	// if tx.RowsAffected == 0 {
	// 	return nil, status.Error(400, "当前订单状态不支持取消")
	// }
	if err != nil {
		return nil, err
	}
	return &order.CommonResp{
		Msg:  "取消订单成功",
		Code: 10501,
	}, nil
}
