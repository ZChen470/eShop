package Catalog

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductLogic) DeleteProduct(req *types.DeleteProductReq) (resp *types.DeleteProductResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.CatalogRpcClient.DeleteProduct(
		l.ctx,
		&catalog.DeleteProductReq{
			ProductId: req.ProductId,
		},
	)
	if err != nil {
		return nil, err
	}
	resp = &types.DeleteProductResp{
		ProductId: r.ProductId,
	}
	return
}
