package helpers

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
)

func NilOrTIme(nullTime sql.NullTime) *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}

	return nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func ValidateStruct(data interface{}) []validator.FieldError {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(data)

	if err != nil {
		var errors []validator.FieldError

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err)
		}

		return errors
	}

	return nil
}
