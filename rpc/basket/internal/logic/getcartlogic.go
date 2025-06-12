package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/basket/basket"
	"github.com/ZChen470/eShop/rpc/basket/internal/svc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartLogic) GetCart(in *basket.GetCartReq) (*basket.GetCartResp, error) {
	// todo: add your logic here and delete this line
	userIdVal := l.ctx.Value("userId")
	if userIdVal == nil {
		return nil, status.Error(500, "用户ID无效")
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%v", userIdVal), 10, 64)
	if err != nil {
		return nil, status.Error(500, "用户ID无效")
	}
	key := fmt.Sprintf("cart:%d", userId)
	itemMap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if err != nil {
		return nil, status.Error(404, "Redis 获取购物车失败")
	}
	var items []*model.CartItem
	for _, v := range itemMap {
		var item model.CartItem
		if err := json.Unmarshal([]byte(v), &item); err != nil {
			continue
		}
		items = append(items, &item)
	}
	// 类型转换
	var basketItems []*basket.CartItem
	for _, item := range items {
		basketItems = append(basketItems, &basket.CartItem{
			ProductId:   int64(item.ProductId),
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}
	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return &basket.GetCartResp{
		Cart: &basket.Cart{
			UserId:     userId,
			Items:      basketItems,
			TotalPrice: total,
		},
	}, nil
}
