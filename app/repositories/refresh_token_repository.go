package repositories

import (
	"gosuper/app/models"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (repo *RefreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return repo.db.Create(refreshToken).Error
}

func (repo *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken

	err := repo.db.Where("token = ?", token).First(&refreshToken).Error

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (repo *RefreshTokenRepository) Delete(refreshToken *models.RefreshToken) error {
	return repo.db.Delete(refreshToken).Error
}

func (repo *RefreshTokenRepository) DeleteAllByUserId(userId string) error {
	return repo.db.Where("user_id = ?", userId).Delete(&models.RefreshToken{}).Error
}

func (repo *RefreshTokenRepository) FindByUserId(userId string) ([]models.RefreshToken, error) {
	var refreshTokens []models.RefreshToken

	err := repo.db.Where("user_id = ?", userId).Find(&refreshTokens).Error

	if err != nil {
		return nil, err
	}

	return refreshTokens, nil
}
