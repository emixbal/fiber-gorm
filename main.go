package main

import (
	"github.com/gofiber/fiber/v2"

	"fiber-gorm/config"
	"fiber-gorm/database"
	"fiber-gorm/routers"
)

func main() {
	app := fiber.New()

	config.InitDB()
	db := config.DB
	database.InitMigration(db)

	routers.Init(app)
	app.Listen(":3000")
}
