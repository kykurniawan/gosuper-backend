package responses

import (
	"time"

	"github.com/google/uuid"
)

type (
	RegisterResponse struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
	}

	LoginResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
)
