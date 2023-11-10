package responses

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserResponse struct {
		ID              uuid.UUID  `json:"id"`
		Name            string     `json:"name"`
		Email           string     `json:"email"`
		EmailVerifiedAt *time.Time `json:"emailVerifiedAt"`
		CreatedAt       *time.Time `json:"createdAt"`
		UpdatedAt       *time.Time `json:"updatedAt"`
		DeletedAt       *time.Time `json:"deletedAt"`
	}

	UserIndexResponse struct {
		Users []UserResponse         `json:"users"`
		Meta  PaginationMetaResponse `json:"meta"`
	}
)
