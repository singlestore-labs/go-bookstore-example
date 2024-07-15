package service

import (
	"bookstore/database"
	"bookstore/model"
	"testing"
)

func ResetBookTable() {
	database.DB.Where("1 = 1").Delete(&model.Book{})
}

func TestCreateBook(t *testing.T) {
	// Arrange
	ResetBookTable()
	book := model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	}

	// Act
	createdBook, err := CreateBook(book)

	// Assert
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	} else if createdBook.ID == 0 {
		t.Errorf("expected non-zero ID, got %d", createdBook.ID)
	} else if createdBook.Title != book.Title {
		t.Errorf("expected %s, got %s", book.Title, createdBook.Title)
	} else if createdBook.Author != book.Author {
		t.Errorf("expected %s, got %s", book.Author, createdBook.Author)
	} else if createdBook.Price != book.Price {
		t.Errorf("expected %f, got %f", book.Price, createdBook.Price)
	} else if createdBook.Genre != book.Genre {
		t.Errorf("expected %s, got %s", book.Genre, createdBook.Genre)
	} else if createdBook.CreatedAt.IsZero() {
		t.Errorf("expected non-zero CreatedAt, got %s", createdBook.CreatedAt)
	} else if createdBook.UpdatedAt.IsZero() {
		t.Errorf("expected non-zero UpdatedAt, got %s", createdBook.UpdatedAt)
	}
}

func TestUpdateBook(t *testing.T) {
	// Arrange
	ResetBookTable()
	book := model.Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Price:  29.99,
		Genre:  "Fiction",
	}

	// Act
	updatedBook, err := UpdateBook(book)

	// Assert
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	} else if updatedBook.ID == 0 {
		t.Errorf("expected non-zero ID, got %d", updatedBook.ID)
	} else if updatedBook.Title != book.Title {
		t.Errorf("expected %s, got %s", book.Title, updatedBook.Title)
	} else if updatedBook.Author != book.Author {
		t.Errorf("expected %s, got %s", book.Author, updatedBook.Author)
	} else if updatedBook.Price != book.Price {
		t.Errorf("expected %f, got %f", book.Price, updatedBook.Price)
	} else if updatedBook.Genre != book.Genre {
		t.Errorf("expected %s, got %s", book.Genre, updatedBook.Genre)
	} else if updatedBook.CreatedAt.IsZero() {
		t.Errorf("expected non-zero CreatedAt, got %s", updatedBook.CreatedAt)
	} else if updatedBook.UpdatedAt.IsZero() {
		t.Errorf("expected non-zero UpdatedAt, got %s", updatedBook.UpdatedAt)
	}
}

func TestDeleteBook(t *testing.T) {
	// Arrange
	ResetBookTable()
	book := model.Book{
		Title:  "To Kill a Mockingbird",
		Author: "Harper Lee",
		Price:  19.99,
		Genre:  "Fiction",
	}
	createdBook, _ := CreateBook(book)

	// Act
	err := DeleteBook(createdBook)

	// Assert
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}

	// Verify the book is deleted
	_, getErr := GetBook(uint(createdBook.ID))
	if getErr == nil {
		t.Errorf("expected error when getting deleted book, got nil")
	}
}

func TestGetBook(t *testing.T) {
	// Arrange
	ResetBookTable()
	book := model.Book{
		Title:  "1984",
		Author: "George Orwell",
		Price:  15.99,
		Genre:  "Science Fiction",
	}
	createdBook, _ := CreateBook(book)

	// Act
	retrievedBook, err := GetBook(uint(createdBook.ID))

	// Assert
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	} else if retrievedBook.ID != createdBook.ID {
		t.Errorf("expected ID %d, got %d", createdBook.ID, retrievedBook.ID)
	} else if retrievedBook.Title != book.Title {
		t.Errorf("expected title %s, got %s", book.Title, retrievedBook.Title)
	} else if retrievedBook.Author != book.Author {
		t.Errorf("expected author %s, got %s", book.Author, retrievedBook.Author)
	} else if retrievedBook.Price != book.Price {
		t.Errorf("expected price %f, got %f", book.Price, retrievedBook.Price)
	} else if retrievedBook.Genre != book.Genre {
		t.Errorf("expected genre %s, got %s", book.Genre, retrievedBook.Genre)
	}
}

func TestGetAllBooks(t *testing.T) {
	// Arrange
	ResetBookTable()
	books := []model.Book{
		{Title: "Pride and Prejudice", Author: "Jane Austen", Price: 12.99, Genre: "Romance"},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", Price: 14.99, Genre: "Fantasy"},
		{Title: "The Catcher in the Rye", Author: "J.D. Salinger", Price: 11.99, Genre: "Fiction"},
	}
	for _, book := range books {
		CreateBook(book)
	}

	// Act
	retrievedBooks := GetAllBooks()

	// Assert
	if len(retrievedBooks) != len(books) {
		t.Errorf("expected %d books, got %d", len(books), len(retrievedBooks))
	}

	for i, book := range retrievedBooks {
		if book.Title != books[i].Title {
			t.Errorf("expected title %s, got %s", books[i].Title, book.Title)
		}
		if book.Author != books[i].Author {
			t.Errorf("expected author %s, got %s", books[i].Author, book.Author)
		}
		if book.Price != books[i].Price {
			t.Errorf("expected price %f, got %f", books[i].Price, book.Price)
		}
		if book.Genre != books[i].Genre {
			t.Errorf("expected genre %s, got %s", books[i].Genre, book.Genre)
		}
	}
}
