package model

import "time"

type Order struct {
	ID        int         `json:"id" gorm:"primaryKey,autoIncrement"`
	Total     float64     `json:"total"`
	Items     []OrderItem `json:"items" gorm:"-"`
	CreatedAt time.Time   `json:"created_at" gorm:"autoCreateTime,precision:6"`
}

type OrderItem struct {
	OrderID  int  `json:"order_id" gorm:"primaryKey"`
	BookID   int  `json:"book_id" gorm:"primaryKey"`
	Book     Book `json:"book" gorm:"-"`
	Quantity int  `json:"quantity"`
}

func (o *Order) CalculateTotal() {
	o.Total = 0
	for _, item := range o.Items {
		o.Total += item.Book.Price * float64(item.Quantity)
	}
}
