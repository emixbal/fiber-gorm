package models

import (
	"fiber-gorm/config"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FethAllBooks() (Response, error) {
	var books []Book
	var res Response

	db := config.DB

	db.Find(&books)

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = books

	return res, nil
}

func CreateABook(name string, email string) (Response, error) {
	var book Book
	var res Response
	db := config.DB

	book.Name = name
	book.Email = email

	if result := db.Create(&book); result.Error != nil {
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = book

	return res, nil
}
