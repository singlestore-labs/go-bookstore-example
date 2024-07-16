package api

import (
	"bookstore/model"
	"bookstore/service"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBook(t *testing.T) {
	// Arrange
	book := model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	}
	reqBody, _ := json.Marshal(book)
	t.Run("Success", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)
		// Assert
		if w.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})
	t.Run("Invalid Book", func(t *testing.T) {
		// Arrange
		invalidBook := model.Book{
			ID: 1,
		}
		reqBody, _ := json.Marshal(invalidBook)
		// Act
		req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)
		// Assert
		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})
	t.Run("Invalid Request", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("POST", "/books", nil)
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)
		// Assert
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestUpdateBook(t *testing.T) {
	// Arrange
	book := model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	}
	book, _ = service.CreateBook(book)
	book.Title = "The Amazing Gatsby"
	reqBody, _ := json.Marshal(book)
	t.Run("Success", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("PUT", fmt.Sprintf("/books/%d", book.ID), bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
	t.Run("Invalid Request", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("PUT", fmt.Sprintf("/books/%d", book.ID), nil)
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
	t.Run("Book Not Found", func(t *testing.T) {
		// Arrange
		notFoundBook := model.Book{
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
			Price:  29.99,
			Genre:  "Fiction",
		}
		reqBody, _ := json.Marshal(notFoundBook)

		// Act
		req := httptest.NewRequest("PUT", "/books/100", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})
}

func TestDeleteBook(t *testing.T) {
	// Arrange
	book := model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	}
	book, _ = service.CreateBook(book)
	t.Run("Success", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/books/%d", book.ID), nil)
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
	t.Run("Invalid Book ID", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("DELETE", "/books/invalid", nil)
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestGetBook(t *testing.T) {
	// Arrange
	book := model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	}
	book, _ = service.CreateBook(book)
	t.Run("Success", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("GET", fmt.Sprintf("/books/%d", book.ID), nil)
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
	t.Run("Book Not Found", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("GET", "/books/100", nil)
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
		}
	})
	t.Run("Invalid Book ID", func(t *testing.T) {
		// Act
		req := httptest.NewRequest("GET", "/books/invalid", nil)
		w := httptest.NewRecorder()
		Router.ServeHTTP(w, req)

		// Assert
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestGetAllBooks(t *testing.T) {
	// Arrange
	book := model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	}
	service.CreateBook(book)
	service.CreateBook(book)
	service.CreateBook(book)

	// Act
	req := httptest.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
