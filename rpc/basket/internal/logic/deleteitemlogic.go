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

type DeleteItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteItemLogic {
	return &DeleteItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteItemLogic) DeleteItem(in *basket.DeleteItemReq) (*basket.CommonResp, error) {
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
	if err := l.svcCtx.Redis.HDel(l.ctx, key, fmt.Sprintf("%d", in.ProductId)).Err(); err != nil {
		return nil, status.Error(500, "删除购物车商品失败")
	}
	return &basket.CommonResp{
		Msg:  "删除购物车商品成功",
		Code: 10303,
	}, nil
}
