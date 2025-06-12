package Basket

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/basket/basket"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemLogic {
	return &UpdateItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateItemLogic) UpdateItem(req *types.UpdateItemReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.BasketRpcClient.AddItem(l.ctx, &basket.AddItemReq{
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	})
	if err != nil {
		return nil, err
	}

	return &types.CommonResp{
		Code: r.Code,
		Msg:  r.Msg,
	}, nil
}
