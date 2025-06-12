package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ZChen470/eShop/rpc/basket/basket"
	"github.com/ZChen470/eShop/rpc/basket/internal/svc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearCartLogic {
	return &ClearCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClearCartLogic) ClearCart(in *basket.ClearCartReq) (*basket.CommonResp, error) {
	// todo: add your logic here and delete this line
	userIdVal := l.ctx.Value("userId")
	if userIdVal == nil {
		return nil, status.Error(500, "用户ID无效")
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%v", userIdVal), 10, 64)
	if err != nil {
		return nil, status.Error(500, "用户ID无效")
	}
	// 构造 Redis key
	key := fmt.Sprintf("cart:%d", userId)
	if err := l.svcCtx.Redis.Del(l.ctx, key).Err(); err != nil {
		return nil, status.Error(500, "清空购物车失败")
	}
	return &basket.CommonResp{
		Msg:  "清空购物车成功",
		Code: 10302,
	}, nil
}
