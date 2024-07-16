package service

import (
	"bookstore/model"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	ResetOrderTable()
	ResetBookTable()
	// Arrange
	gatsby, _ := CreateBook(model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	})
	t.Run("Success", func(t *testing.T) {
		// Act
		order, _ := CreateOrder(model.Order{
			Items: []model.OrderItem{
				{
					BookID:   gatsby.ID,
					Quantity: 2,
				},
			},
		})
		// Assert
		if order.ID == 0 {
			t.Errorf("Expected order ID to be non-zero, got %d", order.ID)
		}
		if order.Items[0].BookID != gatsby.ID {
			t.Errorf("Expected order item book ID to be %d, got %d", gatsby.ID, order.Items[0].BookID)
		} else if order.Items[0].Book.Title != gatsby.Title {
			t.Errorf("Expected order item book title to be %s, got %s", gatsby.Title, order.Items[0].Book.Title)
		} else if order.Total != gatsby.Price*2 {
			t.Errorf("Expected order total to be %f, got %f", gatsby.Price, order.Total)
		} else if order.CreatedAt.IsZero() {
			t.Errorf("Expected order created at to be non-zero, got %v", order.CreatedAt)
		}
	})
	t.Run("Book Not Found", func(t *testing.T) {
		// Act
		_, err := CreateOrder(model.Order{
			Items: []model.OrderItem{
				{
					BookID:   999,
					Quantity: 2,
				},
			},
		})
		// Assert
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
	t.Run("InvalidQuantity", func(t *testing.T) {
		// Act
		_, err := CreateOrder(model.Order{
			Items: []model.OrderItem{
				{
					BookID:   gatsby.ID,
					Quantity: 0,
				},
			},
		})
		// Assert
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
	t.Run("EmptyItems", func(t *testing.T) {
		// Act
		_, err := CreateOrder(model.Order{
			Items: []model.OrderItem{},
		})
		// Assert
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestUpdateOrder(t *testing.T) {
	ResetOrderTable()
	ResetBookTable()
	// Arrange
	gatsby, _ := CreateBook(model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	})
	order, _ := CreateOrder(model.Order{
		Items: []model.OrderItem{
			{
				BookID:   gatsby.ID,
				Quantity: 2,
			},
		},
	})
	t.Run("Success", func(t *testing.T) {
		// Act
		order.Items[0].Quantity = 3
		updatedOrder, _ := UpdateOrder(order)
		// Assert
		if updatedOrder.Total != gatsby.Price*3 {
			t.Errorf("Expected order total to be %f, got %f", gatsby.Price*2, updatedOrder.Total)
		}
	})
	t.Run("InvalidQuantity", func(t *testing.T) {
		// Act
		order.Items[0].Quantity = 0
		_, err := UpdateOrder(order)
		// Assert
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
	t.Run("Book Not Found", func(t *testing.T) {
		// Act
		order.Items[0].BookID = 999
		order.Items[0].Quantity = 2
		_, err := UpdateOrder(order)
		// Assert
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
	t.Run("EmptyItems", func(t *testing.T) {
		// Act
		order.Items = []model.OrderItem{}
		_, err := UpdateOrder(order)
		// Assert
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestDeleteOrder(t *testing.T) {
	ResetOrderTable()
	ResetBookTable()
	// Arrange
	gatsby, _ := CreateBook(model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	})
	order, _ := CreateOrder(model.Order{
		Items: []model.OrderItem{
			{
				BookID:   gatsby.ID,
				Quantity: 2,
			},
		},
	})
	t.Run("Success", func(t *testing.T) {
		// Act
		err := DeleteOrder(order)
		// Assert
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
	t.Run("Error", func(t *testing.T) {
		// Act
		err := DeleteOrder(model.Order{})
		// Assert
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestGetOrder(t *testing.T) {
	// Arrange
	ResetOrderTable()
	ResetBookTable()
	// Arrange
	gatsby, _ := CreateBook(model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	})
	order, _ := CreateOrder(model.Order{
		Items: []model.OrderItem{
			{
				BookID:   gatsby.ID,
				Quantity: 2,
			},
		},
	})
	t.Run("Success", func(t *testing.T) {
		// Act
		retrievedOrder, _ := GetOrder(order.ID)
		// Assert
		if retrievedOrder.ID != order.ID {
			t.Errorf("Expected order ID %d, got %d", order.ID, retrievedOrder.ID)
		} else if retrievedOrder.Items[0].BookID != gatsby.ID {
			t.Errorf("Expected order item book ID to be %d, got %d", gatsby.ID, retrievedOrder.Items[0].BookID)
		}
	})
	t.Run("NotFound", func(t *testing.T) {
		// Act
		_, err := GetOrder(999)
		// Assert
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
func TestGetAllOrders(t *testing.T) {
	// Arrange
	ResetOrderTable()
	ResetBookTable()
	// Arrange
	gatsby, _ := CreateBook(model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	})
	CreateOrder(model.Order{
		Items: []model.OrderItem{
			{
				BookID:   gatsby.ID,
				Quantity: 2,
			},
		},
	})
	t.Run("Success", func(t *testing.T) {
		// Act
		orders := GetAllOrders()
		// Assert
		if len(orders) != 1 {
			t.Errorf("Expected 1 order, got %d", len(orders))
		}
	})
}
