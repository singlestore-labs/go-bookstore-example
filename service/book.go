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

func GetBook(id uint) (model.Book, error) {
	var book model.Book
	result := database.DB.First(&book, "id = ?", id)
	if result.Error != nil {
		return model.Book{}, result.Error
	}
	return book, nil
}

func GetAllBooks() []model.Book {
	var books []model.Book
	database.DB.Find(&books)
	return books
}
