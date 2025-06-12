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

type AddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddItemLogic) AddItem(in *basket.AddItemReq) (*basket.CommonResp, error) {
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
	field := fmt.Sprintf("%d", in.ProductId)

	// 尝试读取已有商品项
	exists, err := l.svcCtx.Redis.HExists(l.ctx, key, field).Result()
	if err != nil {
		return nil, status.Error(500, "Redis 查询失败")
	}
	var item model.CartItem
	if exists {
		// 已存在，取出并更新数量
		val, err := l.svcCtx.Redis.HGet(l.ctx, key, field).Result()
		if err != nil {
			return nil, status.Error(500, "读取购物车项失败")
		}
		if err := json.Unmarshal([]byte(val), &item); err != nil {
			return nil, status.Error(500, "解析购物车项失败")
		}
		// 更新数量
		item.Quantity += in.Quantity
	} else {
		// 不存在，创建新项
		item = model.CartItem{
			ProductId:   uint(in.ProductId),
			ProductName: in.ProductName,
			Price:       float64(in.Price),
			Quantity:    in.Quantity,
		}
	}
	// 存入 Redis
	data, err := json.Marshal(item)
	if err != nil {
		return nil, status.Error(500, "购物车项序列化失败")
	}
	if err := l.svcCtx.Redis.HSet(l.ctx, key, field, data).Err(); err != nil {
		return nil, status.Error(500, "写入 redis 失败")
	}
	return &basket.CommonResp{
		Msg:  "加入购物车成功",
		Code: 10301,
	}, nil
}
