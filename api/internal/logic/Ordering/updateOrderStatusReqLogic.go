package Ordering

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderStatusReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusReqLogic {
	return &UpdateOrderStatusReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderStatusReqLogic) UpdateOrderStatusReq(req *types.UpdateOrderStatusReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.OrderRpcClient.UpdateOrderStatus(
		l.ctx,
		&order.UpdateOrderStatusReq{
			OrderId: req.OrderId,
			Status:  req.Status,
		},
	)
	if err != nil {
		return nil, err
	}
	resp = &types.CommonResp{
		Code: r.Code,
		Msg:  r.Msg,
	}
	return
}
