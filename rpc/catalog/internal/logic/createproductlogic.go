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

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *catalog.CreateProductReq) (*catalog.CreateProductResp, error) {
	// todo: add your logic here and delete this line
	// 开启事务
	tx := l.svcCtx.DB.Begin()
	if in.Name == "" || in.Price <= 0 || in.Stock < 0 || in.Description == "" {
		tx.Rollback()
		return nil, errors.New("参数错误")
	}
	text := in.Name + " " + in.Description
	vector, err := GeneratEmbedding(l.svcCtx.Config.OpenAIKey, text, 384)

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("生成嵌入向量出错 %v", err)
	}

	product := &model.Product{
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
		Stock:       in.Stock,
		Embedding:   vector,
	}

	if err := tx.Create(product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &catalog.CreateProductResp{
		ProductId: int64(product.ProductId),
	}, nil
}
