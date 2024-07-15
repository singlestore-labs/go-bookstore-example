package service

import (
	"bookstore/database"
	"bookstore/model"
)

func CreateBook(book model.Book) (model.Book, error) {
	result := database.DB.Create(&book)
	if result.Error != nil {
		return model.Book{}, result.Error
	}
	return book, nil
}

func UpdateBook(book model.Book) (model.Book, error) {
	return book, nil
}

func DeleteBook(book model.Book) error {
	return nil
}

func GetBook(id uint) (model.Book, error) {
	return model.Book{}, nil
}

func GetAllBooks() []model.Book {
	return []model.Book{}
}
