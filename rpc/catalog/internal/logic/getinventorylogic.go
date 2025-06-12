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

type GetInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInventoryLogic {
	return &GetInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInventoryLogic) GetInventory(in *catalog.GetInventoryReq) (*catalog.GetInventoryResp, error) {
	// todo: add your logic here and delete this line
	if in.ProductId <= 0 {
		return nil, errors.New("商品 ID 错误")
	}
	var product model.Product
	if err := l.svcCtx.DB.Model(&model.Product{}).
		Select("stock").
		Where("product_id = ?", in.ProductId).
		Take(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("商品不存在")
		}
		return nil, err
	}
	return &catalog.GetInventoryResp{
		Stock: product.Stock,
	}, nil
}
