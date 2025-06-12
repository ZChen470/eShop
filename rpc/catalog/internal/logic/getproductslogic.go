package logic

import (
	"context"
	"fmt"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"
	"github.com/ZChen470/eShop/rpc/catalog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductsLogic {
	return &GetProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductsLogic) GetProducts(in *catalog.GetProductsReq) (*catalog.GetProductsResp, error) {
	// todo: add your logic here and delete this line
	page := in.Page
	if page <= 0 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// 查询总数
	var total int64
	if err := l.svcCtx.DB.Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("查询总数失败: %v", err)
	}
	// 查询当前页面商品
	var products []model.Product
	if err := l.svcCtx.DB.Limit(int(pageSize)).
		Offset(int(offset)).
		Find(&products).Error; err != nil {
		return nil, fmt.Errorf("查询当前页面商品失败: %v", err)
	}
	catalogs := make([]*catalog.Product, len(products))
	for i, product := range products {
		catalogs[i] = &catalog.Product{
			ProductId:   int64(product.ProductId),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		}
	}
	return &catalog.GetProductsResp{
		Products: catalogs,
		Total:    int32(total),
	}, nil
}
