package Ordering

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOrderLogic) CancelOrder(req *types.CancelOrderReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.OrderRpcClient.CancelOrder(l.ctx, &order.CancelOrderReq{
		OrderId: req.OrderId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CommonResp{
		Code: r.Code,
		Msg:  r.Msg,
	}, nil
}
