package logic

const (
	OrderStatusPending        = "pending"         // 未支付
	OrderStatusStockConfirmed = "stock_confirmed" // 库存已确认
	OrderStatusUnderstock     = "understock"      // 库存不足
	OrderStatusPaid           = "paid"            // 已支付
	OrderStatusCancelled      = "cancelled"       // 已取消
	OrderStatusShipped        = "shipped"         // 已发货
	OrderStatusCompleted      = "completed"       // 已完成
)

/*
购物车结算->订单生成（pending）->库存检查 -> stock_confirmed ->订单支付（paid）->订单发货（shipped）->订单完成（completed）
订单取消（cancelled）
）
*/
