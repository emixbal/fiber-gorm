package models

import (
	"fiber-gorm/config"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Author string `json:"author"`
	Name   string `json:"name"`
	NoISBN string `json:"no_isbn"`
}

func FethAllBooks() (Response, error) {
	var books []Book
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&books); result.Error != nil {
		fmt.Print("error FethAllBooks")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = books

	return res, nil
}

func CreateABook(book *Book) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&book); result.Error != nil {
		fmt.Print("error CreateABook")
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
