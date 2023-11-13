package exception

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *fiber.Error:
		var fiberError *fiber.Error
		errors.As(err, &fiberError)

		return c.Status(fiberError.Code).JSON(fiber.Map{
			"message": http.StatusText(fiberError.Code),
			"data":    nil,
			"error":   e.Error(),
		})
	case *ValidationError:
		var validationError *ValidationError
		errors.As(err, &validationError)

		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": http.StatusText(http.StatusUnprocessableEntity),
			"data":    nil,
			"error": fiber.Map{
				"fields": validationError.GetFieldErrors(),
				"old":    validationError.Old,
			},
		})
	default:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
			"error":   e.Error(),
		})
	}
}
