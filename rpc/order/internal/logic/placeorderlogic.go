package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/basket/basketclient"
	"github.com/ZChen470/eShop/rpc/order/internal/svc"
	"github.com/ZChen470/eShop/rpc/order/order"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceOrderLogic {
	return &PlaceOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlaceOrderLogic) PlaceOrder(in *order.PlaceOrderReq) (*order.PlaceOrderResp, error) {
	// 客户端创建订单时，通过购物车数据，创建订单
	userIdVal := l.ctx.Value("userId")
	if userIdVal == nil {
		return nil, status.Error(500, "用户ID无效")
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%v", userIdVal), 10, 64)
	if err != nil {
		return nil, status.Error(500, "用户ID无效")
	}

	// 查询商品信息 & 构造订单项
	var orderItems []model.OrderItem
	var totalAmount float64
	// 构造订单项
	for _, item := range in.Items {
		var product model.Product
		if err := l.svcCtx.DB.First(&product, item.ProductId).Error; err != nil {
			return nil, status.Errorf(404, "商品ID %d 不存在", item.ProductId)
		}
		orderItem := model.OrderItem{
			ProductId: product.ProductId,
			Price:     product.Price,
			Quantity:  item.Quantity,
		}
		orderItems = append(orderItems, orderItem)
		totalAmount += product.Price * float64(item.Quantity)
	}
	// 构造订单
	ord := model.Order{
		UserId:      uint(userId),
		Status:      OrderStatusPending,
		TotalAmount: totalAmount,
		Items:       orderItems,
	}
	// 开启事务保存订单及订单项
	if err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&ord).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, status.Error(500, "创建订单失败")
	}

	// 构造创建订单事件消息
	type product struct {
		ProductId uint `json:"productId"`
		Quantity  int  `json:"quantity"`
	}

	products := []product{}
	for _, v := range in.Items {
		products = append(products, product{
			ProductId: uint(v.ProductId),
			Quantity:  int(v.Quantity),
		})
	}
	productVal, err := json.Marshal(products)
	if err != nil {
		return nil, fmt.Errorf("序列化商品信息失败: %v", err)
	}
	// 推送消息给库存服务
	l.svcCtx.KqPusherClient.KPush(l.ctx, strconv.Itoa(int(ord.OrderId)), string(productVal))
	// 清空购物车
	l.svcCtx.BasketRpcClient.ClearCart(l.ctx, &basketclient.ClearCartReq{})

	return &order.PlaceOrderResp{
		OrderId: int32(ord.OrderId),
	}, nil
}
