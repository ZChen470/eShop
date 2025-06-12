package Basket

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/basket/basket"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartLogic) GetCart() (resp *types.GetCartResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.BasketRpcClient.GetCart(
		l.ctx,
		&basket.GetCartReq{},
	)
	if err != nil {
		return nil, err
	}

	// type Cart struct {
	// 	UserId     int64
	// 	Items      []CartItem
	// 	TotalPrice float64
	// }
	items := make([]types.CartItem, 0)
	for _, item := range r.Cart.Items {
		items = append(items, types.CartItem{
			ProductId:   item.ProductId,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}
	cart := types.Cart{
		UserId:     r.Cart.UserId,
		Items:      items,
		TotalPrice: r.Cart.TotalPrice,
	}
	return &types.GetCartResp{
		Cart: cart,
	}, nil
}
