package middlewares

import (
	"gosuper/app/models"
	"gosuper/app/services"

	"github.com/gofiber/fiber/v2"
)

func Verified(authService *services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authenticatedUser := c.Locals("user").(*models.User)

		if !authenticatedUser.EmailVerifiedAt.Valid {
			return fiber.NewError(fiber.StatusForbidden, "Email is not verified")
		}

		return c.Next()
	}
}
