package requests

import "github.com/go-playground/validator/v10"

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
)

func (request *RegisterRequest) Validate() []validator.FieldError {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(request)

	if err != nil {
		var errors []validator.FieldError

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err)
		}

		return errors
	}

	return nil
}

func (request *LoginRequest) Validate() []validator.FieldError {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(request)

	if err != nil {
		var errors []validator.FieldError

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err)
		}

		return errors
	}

	return nil
}
