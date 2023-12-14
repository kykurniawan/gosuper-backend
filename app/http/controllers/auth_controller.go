package controllers

import (
	"gosuper/app/constants"
	"gosuper/app/exception"
	"gosuper/app/http/requests"
	"gosuper/app/http/responses"
	"gosuper/app/models"
	"gosuper/app/services"
	"gosuper/config"

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
	c.Accepts(constants.JsonContentType)

	loginRequest := new(requests.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrorRequestBodyNotValid)
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
		"data": responses.TokenResponse{
			AccessToken:           accessToken,
			AccessTokenExpiresIn:  config.Token.AccessTokenExpiresIn,
			RefreshToken:          refreshToken,
			RefreshTokenExpiresIn: config.Token.RefeshTokenExpiresIn,
		},
		"errors": nil,
	})
}

func (controller *AuthController) Logout(c *fiber.Ctx) error {
	c.Accepts(constants.JsonContentType)

	logoutRequest := new(requests.LogoutRequest)

	if err := c.BodyParser(logoutRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrorRequestBodyNotValid)
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
	c.Accepts()

	registerRequest := new(requests.RegisterRequest)

	if err := c.BodyParser(registerRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrorRequestBodyNotValid)
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
	c.Accepts(constants.JsonContentType)

	refreshTokenRequest := new(requests.RefreshTokenRequest)

	if err := c.BodyParser(refreshTokenRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrorRequestBodyNotValid)
	}

	if err := refreshTokenRequest.Validate(); err != nil {
		return exception.NewValidationError(err, refreshTokenRequest)
	}

	accessToken, refreshToken, err := controller.authService.RefreshToken(refreshTokenRequest)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Refresh token success!",
		"data": responses.TokenResponse{
			AccessToken:           accessToken,
			AccessTokenExpiresIn:  config.Token.AccessTokenExpiresIn,
			RefreshToken:          refreshToken,
			RefreshTokenExpiresIn: config.Token.RefeshTokenExpiresIn,
		},
		"errors": nil,
	})
}

func (controller *AuthController) User(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	loggedUser := responses.LoggedUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(fiber.Map{
		"message": "Get user success!",
		"data":    loggedUser,
		"errors":  nil,
	})
}

func (controller *AuthController) ForgotPassword(c *fiber.Ctx) error {
	c.Accepts(constants.JsonContentType)

	forgotPasswordRequest := new(requests.ForgotPasswordRequest)

	if err := c.BodyParser(forgotPasswordRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrorRequestBodyNotValid)
	}

	if err := forgotPasswordRequest.Validate(); err != nil {
		return exception.NewValidationError(err, forgotPasswordRequest)
	}

	if err := controller.authService.ForgotPassword(forgotPasswordRequest); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "We have sent you an otp to your email!",
		"data":    nil,
		"errors":  nil,
	})
}

func (controller *AuthController) ResetPassword(c *fiber.Ctx) error {
	c.Accepts(constants.JsonContentType)

	resetPasswordRequest := new(requests.ResetPasswordRequest)

	if err := c.BodyParser(resetPasswordRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrorRequestBodyNotValid)
	}

	if err := resetPasswordRequest.Validate(); err != nil {
		return exception.NewValidationError(err, resetPasswordRequest)
	}

	if err := controller.authService.ResetPassword(resetPasswordRequest); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Reset password success!",
		"data":    nil,
		"errors":  nil,
	})
}

func (controller *AuthController) ResendEmailVerification(c *fiber.Ctx) error {
	c.Accepts(constants.JsonContentType)

	authenticatedUser := c.Locals("user").(*models.User)

	if err := controller.authService.ResendEmailVerification(authenticatedUser); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "We have sent you an otp to your email!",
		"data":    nil,
		"errors":  nil,
	})
}

func (controller *AuthController) VerifyEmail(c *fiber.Ctx) error {
	c.Accepts(constants.JsonContentType)

	verifyEmailRequest := new(requests.VerifyEmailRequest)

	if err := c.BodyParser(verifyEmailRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrorRequestBodyNotValid)
	}

	if err := verifyEmailRequest.Validate(); err != nil {
		return exception.NewValidationError(err, verifyEmailRequest)
	}

	user, err := controller.authService.VerifyEmail(verifyEmailRequest)

	if err != nil {
		return err
	}

	c.Locals("user", user)

	return c.JSON(fiber.Map{
		"message": "Email has been verified!",
		"data":    nil,
		"errors":  nil,
	})
}
