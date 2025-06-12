package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"
	"github.com/ZChen470/eShop/rpc/catalog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *catalog.UpdateProductReq) (*catalog.UpdateProductResp, error) {
	// todo: add your logic here and delete this line
	tx := l.svcCtx.DB.Begin()
	if in.Name == "" || in.Price <= 0 || in.Stock < 0 || in.Description == "" {
		tx.Rollback()
		return nil, errors.New("参数错误")
	}
	text := in.Name + " " + in.Description
	vector, err := GeneratEmbedding(l.svcCtx.Config.OpenAIKey, text, 384)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("生成嵌入向量出错 %f", err)
	}
	if err := tx.Model(&model.Product{}).Where("product_i = ?", in.ProductId).Updates(
		map[string]interface{}{
			"name":        in.Name,
			"description": in.Description,
			"price":       in.Price,
			"stock":       in.Stock,
			"embedding":   vector,
		},
	).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("更新商品信息失败：%v", err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("提交事务失败：%v", err)
	}
	return &catalog.UpdateProductResp{
		ProductId: in.ProductId,
	}, nil
}
