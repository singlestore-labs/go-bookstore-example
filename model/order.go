package model

import "time"

type Order struct {
	ID        int         `json:"id" gorm:"primaryKey,autoIncrement"`
	Total     float64     `json:"total"`
	Items     []OrderItem `json:"items" gorm:"-"`
	CreatedAt time.Time   `json:"created_at" gorm:"autoCreateTime,precision:6"`
}

type OrderItem struct {
	OrderID  int `json:"order_id" gorm:"primaryKey"`
	BookID   int `json:"book_id" gorm:"primaryKey"`
	Quantity int `json:"quantity"`
}
