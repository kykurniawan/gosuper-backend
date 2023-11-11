package repositories

import (
	"gosuper/app/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	Database *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		Database: db,
	}
}

func (repo *RefreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return repo.Database.Create(refreshToken).Error
}
