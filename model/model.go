package model

import "time"

const (
	OrderStatusPending   = "pending"
	OrderStatusCompleted = "completed"
	OrderStatusCancelled = "cancelled"
	OrderStatusShipped   = "shipped"
)

// Product 商品表
type Product struct {
	ProductId   uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"type:varchar(100);not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int32   `gorm:"not null" json:"stock"`
	Embedding   Vector  `gorm:"type:vector(384)"`
}

// Order 订单表
type Order struct {
	OrderId     uint        `gorm:"primaryKey" json:"order_id"`
	UserId      uint        `gorm:"index;not null" json:"user_id"`
	Status      string      `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TotalAmount float64     `gorm:"not null" json:"total_amount"`
	Items       []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt   time.Time   `json:"created_at"` // GORM 会自动维护
	UpdatedAt   time.Time   `json:"updated_at"` // GORM 会自动维护
}

// OrderItem 订单项
type OrderItem struct {
	ItemId    uint    `gorm:"primaryKey" json:"item_id"`
	OrderId   uint    `gorm:"index;not null" json:"order_id"`
	ProductId uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int32   `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
}

// User 用户表
type User struct {
	UserId   uint   `gorm:"primaryKey" json:"user_id"`
	Nickname string `gorm:"uniqueIndex;type:varchar(100);not null" json:"nickname"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
}

// Redis
// CartItem 购物车项
type CartItem struct {
	ProductId   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int32   `json:"quantity"`
}

// Cart 购物车
type Cart struct {
	UserId uint       `json:"user_id"`
	Items  []CartItem `json:"items"`
}
