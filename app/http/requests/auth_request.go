package requests

import (
	"gosuper/app/helpers"

	"github.com/go-playground/validator/v10"
)

type BaseRequest struct{}

func (base *BaseRequest) Validate() []validator.FieldError {
	return helpers.ValidateStruct(base)
}

type RegisterRequest struct {
	BaseRequest
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	BaseRequest
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LogoutRequest struct {
	BaseRequest
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshTokenRequest struct {
	BaseRequest
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type ForgotPasswordRequest struct {
	BaseRequest
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	BaseRequest
	Otp      string `json:"otp" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
