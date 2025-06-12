package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/catalog/catalog"
	"github.com/ZChen470/eShop/rpc/catalog/internal/svc"
	"github.com/lib/pq"

	"github.com/zeromicro/go-zero/core/logx"
)

type SemanticSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSemanticSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SemanticSearchLogic {
	return &SemanticSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SemanticSearchLogic) SemanticSearch(in *catalog.SemanticSearchReq) (*catalog.SemanticSearchResp, error) {
	// todo: add your logic here and delete this line
	if strings.TrimSpace(in.Query) == "" {
		return nil, errors.New("搜索关键词不能为空")
	}
	// 生成查询向量
	queryVector, err := GeneratEmbedding(l.svcCtx.Config.OpenAIKey, in.Query, 384)
	if err != nil {
		return nil, fmt.Errorf("生成查询嵌入向量失败")
	}
	// 使用 gpvector 进行相似度查询，余弦相似度
	var products []model.Product
	if err := l.svcCtx.DB.Raw(`
		SELECT * FROM products
		ORDER BY embedding <#> ?::vector
		LIMIT 5
	`, pq.Array(queryVector)).Scan(&products).Error; err != nil {
		return nil, fmt.Errorf("数据库语义搜索失败：%v", err)
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

	return &catalog.SemanticSearchResp{
		Products: catalogs,
	}, nil
}
