package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/order/internal/svc"
	"github.com/ZChen470/eShop/rpc/order/order"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderListLogic) GetOrderList(in *order.GetOrderListReq) (*order.GetOrderListResp, error) {
	// todo: add your logic here and delete this line
	userIdVal := l.ctx.Value("userId")
	if userIdVal == nil {
		return nil, status.Error(500, "用户ID无效")
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%v", userIdVal), 10, 64)
	if err != nil {
		return nil, status.Error(500, "用户ID无效")
	}
	// 查询订单及订单项和商品信息
	var orders []model.Order
	if err := l.svcCtx.DB.Preload("Items.Product").
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Find(&orders).Error; err != nil {
		return nil, status.Error(500, "查询订单失败")
	}
	var respOrders []*order.OrderProfile
	for _, o := range orders {
		var productNames []string
		for _, item := range o.Items {
			productNames = append(productNames, item.Product.Name)
		}
		respOrders = append(respOrders, &order.OrderProfile{
			OrderId:     int32(o.OrderId),
			UserId:      int32(o.OrderId),
			Status:      o.Status,
			TotalAmount: o.TotalAmount,
			ProductName: productNames,
		})
	}
	return &order.GetOrderListResp{
		Orders: respOrders,
	}, nil
}
