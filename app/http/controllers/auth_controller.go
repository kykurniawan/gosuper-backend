package controllers

import (
	"gosuper/app/exception"
	"gosuper/app/http/requests"
	"gosuper/app/http/responses"
	"gosuper/app/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	c.Accepts("application/json")

	loginRequest := new(requests.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Request body is not valid!")
	}

	if err := loginRequest.Validate(); err != nil {
		return exception.NewValidationError(err, loginRequest)
	}

	accessToken, refreshToken, err := controller.authService.Login(loginRequest)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Login success!",
		"data": responses.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		"errors": nil,
	})
}

func (controller *AuthController) Logout(c *fiber.Ctx) error {
	c.Accepts("application/json")

	logoutRequest := new(requests.LogoutRequest)

	if err := c.BodyParser(logoutRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Request body is not valid!")
	}

	if err := logoutRequest.Validate(); err != nil {
		return exception.NewValidationError(err, logoutRequest)
	}

	if err := controller.authService.Logout(logoutRequest); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Logout success!",
		"data":    nil,
		"errors":  nil,
	})
}

func (controller *AuthController) Register(c *fiber.Ctx) error {
	c.Accepts("application/json")

	registerRequest := new(requests.RegisterRequest)

	if err := c.BodyParser(registerRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Request body is not valid!")
	}

	if err := registerRequest.Validate(); err != nil {
		return exception.NewValidationError(err, registerRequest)
	}

	user, err := controller.authService.Register(registerRequest)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Register success!",
		"data": responses.RegisterResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time,
		},
		"errors": nil,
	})
}

func (controller *AuthController) Refresh(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Refresh success!",
		"data":    nil,
		"errors":  nil,
	})
}

func (controller *AuthController) User(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Get user success!",
		"data":    nil,
		"errors":  nil,
	})
}
