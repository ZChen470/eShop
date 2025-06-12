package logic

import (
	"context"
	"errors"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"
	"github.com/ZChen470/eShop/rpc/catalog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *catalog.DeleteProductReq) (*catalog.DeleteProductResp, error) {
	// todo: add your logic here and delete this line
	if in.ProductId <= 0 {
		return nil, errors.New("商品 ID 错误")
	}
	// tx := l.svcCtx.DB.Begin() 删除暂时不需要事务
	if err := l.svcCtx.DB.Where("product_id = ?", in.ProductId).Delete(&model.Product{}).Error; err != nil {
		return nil, err
	}
	return &catalog.DeleteProductResp{
		ProductId: in.ProductId,
	}, nil
}
