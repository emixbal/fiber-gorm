package middlewares

import (
	"fiber-gorm/config"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthenticated(c *fiber.Ctx) error {
	raw_token := c.Request().Header.Peek("Authorization")
	tokenString := string(raw_token)
	conf := config.GetConfig()

	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(
			map[string]string{
				"message": "Unauthorized, need access token to access this API route!",
			},
		)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		key := []byte(conf.JWT_SECRET)
		return key, nil
	})
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusUnauthorized).JSON(
			map[string]string{
				"message": "Unauthorized, access token is invalid!",
			},
		)
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return c.Next()
	}
	return c.Status(http.StatusUnauthorized).JSON(
		map[string]string{
			"message": "Unauthorized, access token is invalid!",
		},
	)
}
