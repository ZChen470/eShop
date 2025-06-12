package Ordering

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(req *types.GetOrderDetailReq) (resp *types.GetOrderDetailResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.OrderRpcClient.GetOrderDetail(l.ctx, &order.GetOrderDetailReq{
		OrderId: req.OrderId,
	})
	if err != nil {
		return nil, err
	}
	items := make([]types.OrderItem, 0)
	for _, item := range r.Order.Items {
		items = append(items, types.OrderItem{
			ProductId:   item.ProductId,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}
	return &types.GetOrderDetailResp{
		Order: types.Order{
			OrderId:     r.Order.OrderId,
			UserId:      r.Order.UserId,
			TotalAmount: r.Order.TotalAmount,
			CreateAt:    r.Order.CreateAt,
			UpdateAt:    r.Order.UpdateAt,
			Status:      r.Order.Status,
			Items:       items,
		},
	}, nil
}
