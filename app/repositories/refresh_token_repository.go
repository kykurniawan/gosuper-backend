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

func (repo *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken

	err := repo.Database.Where("token = ?", token).First(&refreshToken).Error

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (repo *RefreshTokenRepository) Delete(refreshToken *models.RefreshToken) error {
	return repo.Database.Delete(refreshToken).Error
}

func (repo *RefreshTokenRepository) DeleteAllByUserId(userId string) error {
	return repo.Database.Where("user_id = ?", userId).Delete(&models.RefreshToken{}).Error
}
