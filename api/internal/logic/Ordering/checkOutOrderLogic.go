package Ordering

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/order/ordering"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckOutOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckOutOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckOutOrderLogic {
	return &CheckOutOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckOutOrderLogic) CheckOutOrder(req *types.CheckOutOrderReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.OrderRpcClient.CheckOutOrder(l.ctx, &ordering.CheckOutOrderReq{
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
