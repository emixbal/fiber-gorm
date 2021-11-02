package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ExampleMiddleware(c *fiber.Ctx) error {
	fmt.Println("ExampleMiddleware hitted")
	return c.Next()
}

func FetchUserMiddleware(c *fiber.Ctx) error {
	fmt.Println("FetchUserMiddleware hitted")
	return c.Next()
}
