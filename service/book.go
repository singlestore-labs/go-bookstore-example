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
	result := database.DB.Model(&book).Updates(book)
	if result.Error != nil {
		return model.Book{}, result.Error
	}
	return book, nil
}

func DeleteBook(book model.Book) error {
	result := database.DB.Delete(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetBook(id int) model.Book {
	book := model.Book{}
	database.DB.Find(&book, "id = ?", id)
	return book
}

func GetAllBooks() []model.Book {
	books := []model.Book{}
	database.DB.Find(&books)
	return books
}
