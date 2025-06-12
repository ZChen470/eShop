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

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(in *order.GetOrderDetailReq) (*order.GetOrderDetailResp, error) {
	// todo: add your logic here and delete this line
	userIdVal := l.ctx.Value("userId")
	if userIdVal == nil {
		return nil, status.Error(500, "用户ID无效")
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%v", userIdVal), 10, 64)
	if err != nil {
		return nil, status.Error(500, "用户ID无效")
	}
	o := new(model.Order)
	if err := l.svcCtx.DB.
		Preload("Items.Product").
		Where("order_id = ? AND user_id = ?", in.OrderId, userId).
		Find(o).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(404, "订单不存在")
		}
		return nil, status.Error(500, "查询订单失败")
	}
	var items []*order.OrderItem
	for _, v := range o.Items {
		items = append(items, &order.OrderItem{
			ProductId:   int32(v.ProductId),
			ProductName: v.Product.Name,
			Price:       v.Price,
			Quantity:    v.Quantity,
		})
	}
	return &order.GetOrderDetailResp{
		Order: &order.Order{
			OrderId:     int32(o.OrderId),
			UserId:      int32(userId),
			TotalAmount: o.TotalAmount,
			Items:       items,
			Status:      o.Status,
			CreateAt:    o.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:    o.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
