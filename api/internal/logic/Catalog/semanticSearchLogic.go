package Catalog

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

type SemanticSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSemanticSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SemanticSearchLogic {
	return &SemanticSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SemanticSearchLogic) SemanticSearch(req *types.SemanticSearchReq) (resp *types.SemanticSearchResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.CatalogRpcClient.SemanticSearch(l.ctx, &catalog.SemanticSearchReq{
		Query: req.Query,
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
	return &types.SemanticSearchResp{
		Products: products,
	}, nil
}
