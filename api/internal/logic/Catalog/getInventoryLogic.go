package Catalog

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInventoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInventoryLogic {
	return &GetInventoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInventoryLogic) GetInventory(req *types.GetInventoryReq) (resp *types.GetInventoryResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.CatalogRpcClient.GetInventory(l.ctx, &catalog.GetInventoryReq{
		ProductId: req.ProductId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetInventoryResp{
		ProductId: r.ProductId,
		Stock:     r.Stock,
	}
	return
}
