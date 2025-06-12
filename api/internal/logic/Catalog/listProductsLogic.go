package Catalog

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductsLogic {
	return &ListProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProductsLogic) ListProducts() (resp *types.GetProductsResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.CatalogRpcClient.GetProducts(l.ctx, &catalog.GetProductsReq{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, len(r.Products))
	for i, product := range r.Products {
		products[i] = types.Product{
			ProductId:   product.ProductId,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		}
	}
	resp = &types.GetProductsResp{
		Products: products,
		Total:    r.Total,
	}
	return
}
