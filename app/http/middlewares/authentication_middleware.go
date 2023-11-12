package middlewares

import (
	"gosuper/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(authService *services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token string

		authorizationHeader := c.Get("Authorization")

		if authorizationHeader == "" {
			return fiber.ErrUnauthorized
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			return fiber.ErrUnauthorized
		}

		token = strings.TrimPrefix(authorizationHeader, "Bearer ")

		user, err := authService.ValidateAccessToken(token)

		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		c.Locals("user", user)

		return c.Next()
	}
}
