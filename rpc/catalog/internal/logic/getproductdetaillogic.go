package logic

import (
	"context"
	"errors"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"
	"github.com/ZChen470/eShop/rpc/catalog/internal/svc"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductDetailLogic {
	return &GetProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductDetailLogic) GetProductDetail(in *catalog.GetProductDetailReq) (*catalog.GetProductDetailResp, error) {
	// todo: add your logic here and delete this line
	if in.ProductId <= 0 {
		return nil, errors.New("商品 ID 错误")
	}
	var product model.Product
	if err := l.svcCtx.DB.
		Where("product_id = ?", in.ProductId).
		Take(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("商品不存在")
		}
		return nil, err
	}
	return &catalog.GetProductDetailResp{
		Product: &catalog.Product{
			ProductId:   int64(product.ProductId),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		},
	}, nil
}
