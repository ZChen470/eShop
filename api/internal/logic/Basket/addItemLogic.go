package Basket

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/basket/basket"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddItemLogic) AddItem(req *types.AddItemReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.BasketRpcClient.AddItem(
		l.ctx,
		&basket.AddItemReq{
			ProductId:   req.ProductId,
			ProductName: req.ProductName,
			Price:       req.Price,
			Quantity:    req.Quantity,
		},
	)
	if err != nil {
		return nil, err
	}
	return &types.CommonResp{
		Code: r.Code,
		Msg:  r.Msg,
	}, nil
}
