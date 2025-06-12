package Catalog

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductDetailLogic {
	return &GetProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductDetailLogic) GetProductDetail(req *types.GetProductDetailReq) (resp *types.GetProductDetailResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.CatalogRpcClient.GetProductDetail(
		l.ctx,
		&catalog.GetProductDetailReq{
			ProductId: req.ProductId,
		},
	)
	if err != nil {
		return nil, err
	}
	resp = &types.GetProductDetailResp{
		Product: types.Product{
			ProductId:   r.Product.ProductId,
			Name:        r.Product.Name,
			Price:       r.Product.Price,
			Stock:       r.Product.Stock,
			Description: r.Product.Description,
		},
	}
	return
}
