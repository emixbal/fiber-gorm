package routers

import (
	"fiber-gorm/controllers"

	"github.com/gofiber/fiber/v2"
)

func Book(app *fiber.App) {
	user := app.Group("/books")
	user.Get("/",
		controllers.FetchAllBooks,
	)

	user.Post("/",
		controllers.CreateBook,
	)
}
