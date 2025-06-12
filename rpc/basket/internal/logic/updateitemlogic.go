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

type UpdateItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemLogic {
	return &UpdateItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateItemLogic) UpdateItem(in *basket.UpdateItemReq) (*basket.CommonResp, error) {
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
	// 尝试查询商品项信息
	exists, err := l.svcCtx.Redis.HExists(l.ctx, key, field).Result()
	if err != nil {
		return nil, status.Error(404, "Redis 查询失败")
	}
	var item model.CartItem
	if exists {
		val, err := l.svcCtx.Redis.HGet(l.ctx, key, field).Result()
		if err != nil {
			return nil, status.Error(500, "读取购物车项失败")
		}
		if err := json.Unmarshal([]byte(val), &item); err != nil {
			return nil, status.Error(500, "解析购物车项失败")
		}
		item.Quantity = in.Quantity
	} else {
		return nil, status.Error(404, "商品项不存在")
	}
	data, err := json.Marshal(item)
	if err != nil {
		return nil, status.Error(500, "序列化购物车项失败")
	}
	if err := l.svcCtx.Redis.HSet(l.ctx, key, field, data).Err(); err != nil {
		return nil, status.Error(500, "写入 Redis 失败")
	}
	return &basket.CommonResp{
		Msg:  "更新购物车商品项成功",
		Code: 10304,
	}, nil
}
