package Ordering

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/order/ordering"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListLogic) GetOrderList() (resp *types.GetOrderListResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.OrderRpcClient.GetOrderList(l.ctx, &ordering.GetOrderListReq{})
	if err != nil {
		return nil, err
	}

	orders := make([]types.OrderProfile, 0)
	for _, order := range r.Orders {
		orders = append(orders, types.OrderProfile{
			OrderId:     order.OrderId,
			UserId:      order.UserId,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
			ProductName: order.ProductName,
		})
	}
	resp = &types.GetOrderListResp{
		Orders: orders,
	}
	return
}
