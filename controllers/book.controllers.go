package controllers

import (
	"fiber-gorm/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func FetchAllBooks(c *fiber.Ctx) error {
	result, err := models.FethAllBooks()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(result)
}

func CreateBook(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	if name == "" {
		return c.Status(402).SendString("Error while save book")
	}
	if email == "" {
		return c.Status(402).SendString("Error while save book")
	}

	result, err := models.CreateABook(name, email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(result)
}
