package routers

import (
	"fiber-gorm/controllers"
	"fiber-gorm/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Book(app *fiber.App) {
	user := app.Group("/books")

	user.Get("/", middlewares.ExampleMiddleware, controllers.FetchAllBooks) // contoh menggunakan middleware
	user.Post("/", controllers.CreateBook)
}
