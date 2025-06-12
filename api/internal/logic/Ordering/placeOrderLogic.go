package Ordering

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceOrderLogic {
	return &PlaceOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceOrderLogic) PlaceOrder(req *types.PlaceOrderReq) (resp *types.PlaceOrderResp, err error) {
	// todo: add your logic here and delete this line
	items := make([]*order.OrderItem, 0)
	for _, item := range req.Items {
		items = append(items, &order.OrderItem{
			ProductId:   item.ProductId,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}
	r, err := l.svcCtx.OrderRpcClient.PlaceOrder(
		l.ctx,
		&order.PlaceOrderReq{
			Items: items,
		},
	)
	if err != nil {
		return
	}
	resp = &types.PlaceOrderResp{
		OrderId: r.OrderId,
	}
	return
}
