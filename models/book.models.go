package models

import (
	"fiber-gorm/db"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name string `json:"name"`
}

func FethAllBooks() (Response, error) {
	var books []Book
	var res Response

	db := db.CreateCon()

	db.Find(&books)

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = books

	return res, nil
}

func CreateABook(name string) (Response, error) {
	var book Book
	var res Response
	db := db.CreateCon()

	book.Name = name

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
