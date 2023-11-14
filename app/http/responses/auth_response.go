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

	TokenResponse struct {
		AccessToken           string `json:"accessToken"`
		AccessTokenExpiresIn  int    `json:"accessTokenExpiresIn"`
		RefreshToken          string `json:"refreshToken"`
		RefreshTokenExpiresIn int    `json:"refreshTokenExpiresIn"`
	}

	LoggedUserResponse struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
	}
)
