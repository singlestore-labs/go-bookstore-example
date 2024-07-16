package service

import (
	"bookstore/database"
	"bookstore/model"
	"errors"
)

func CreateOrder(order model.Order) (model.Order, error) {
	if len(order.Items) == 0 {
		return model.Order{}, errors.New("order items are required")
	}
	for i, item := range order.Items {
		if item.Quantity <= 0 {
			return model.Order{}, errors.New("quantity must be greater than 0")
		}
		book, err := GetBook(item.BookID)
		if err != nil {
			return model.Order{}, err
		}
		order.Items[i].Book = book
	}
	order.CalculateTotal()
	database.DB.Create(&order)
	for i := range order.Items {
		order.Items[i].OrderID = order.ID
	}
	database.DB.Create(&order.Items)
	return order, nil
}

func UpdateOrder(order model.Order) (model.Order, error) {
	if len(order.Items) == 0 {
		return model.Order{}, errors.New("order items are required")
	}
	for i, item := range order.Items {
		if item.Quantity <= 0 {
			return model.Order{}, errors.New("quantity must be greater than 0")
		}
		book, err := GetBook(item.BookID)
		if err != nil {
			return model.Order{}, err
		}
		order.Items[i].Book = book
	}
	order.CalculateTotal()
	database.DB.Model(&order).Updates(order)
	for i := range order.Items {
		order.Items[i].OrderID = order.ID
	}
	database.DB.Where("order_id = ?", order.ID).Delete(&model.OrderItem{})
	database.DB.Model(&model.OrderItem{}).Updates(order.Items)
	return order, nil
}

func DeleteOrder(order model.Order) error {
	result := database.DB.Delete(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetOrder(id int) (model.Order, error) {
	var order model.Order
	result := database.DB.Where("id = ?", id).First(&order)
	if result.Error != nil {
		return model.Order{}, result.Error
	}
	database.DB.Model(&model.OrderItem{}).Where("order_id = ?", order.ID).Find(&order.Items)
	return order, nil
}

func GetAllOrders() []model.Order {
	var orders []model.Order
	database.DB.Find(&orders)
	for i := range orders {
		database.DB.Model(&model.OrderItem{}).Where("order_id = ?", orders[i].ID).Find(&orders[i].Items)
	}
	return orders
}
