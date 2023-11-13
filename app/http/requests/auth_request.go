package requests

import (
	"gosuper/app/helpers"

	"github.com/go-playground/validator/v10"
)

type (
	RegisterRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LogoutRequest struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}

	ForgotPasswordRequest struct {
		Email string `json:"email" validate:"required,email"`
	}

	ResetPasswordRequest struct {
		Otp      string `json:"otp" validate:"required"`
		Password string `json:"password" validate:"required,min=8"`
	}
)

func (request *RegisterRequest) Validate() []validator.FieldError {
	return helpers.ValidateStruct(request)
}

func (request *LoginRequest) Validate() []validator.FieldError {
	return helpers.ValidateStruct(request)
}

func (request *LogoutRequest) Validate() []validator.FieldError {
	return helpers.ValidateStruct(request)
}

func (request *RefreshTokenRequest) Validate() []validator.FieldError {
	return helpers.ValidateStruct(request)
}

func (request *ForgotPasswordRequest) Validate() []validator.FieldError {
	return helpers.ValidateStruct(request)
}

func (request *ResetPasswordRequest) Validate() []validator.FieldError {
	return helpers.ValidateStruct(request)
}
