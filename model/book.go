package model

import "time"

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey,autoIncrement"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Genre     string    `json:"genre"`
	Price     float64   `json:"price"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime,precision:6"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime,precision:6"`
}
