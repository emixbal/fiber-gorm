package main

import (
	"github.com/gofiber/fiber/v2"

	"fiber-gorm/db"
	"fiber-gorm/routers"
)

func main() {
	app := fiber.New()

	routers.Init(app)

	// create migration
	// db.CreateCon().AutoMigrate(&models.Book{})

	db.Init()
	app.Listen(":3000")
}
