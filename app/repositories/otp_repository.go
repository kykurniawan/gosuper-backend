package repositories

import (
	"gosuper/app/models"

	"gorm.io/gorm"
)

type OtpRepository struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) *OtpRepository {
	return &OtpRepository{
		db: db,
	}
}

func (repo *OtpRepository) Create(otp *models.Otp) error {
	return repo.db.Create(otp).Error
}

func (repo *OtpRepository) Update(otp *models.Otp) error {
	return repo.db.Save(otp).Error
}

func (repo *OtpRepository) FindByCodeAndType(code string, otpType string) (*models.Otp, error) {
	var otp models.Otp

	err := repo.db.Where("code = ? AND otp_type = ?", code, otpType).First(&otp).Error

	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (repo *OtpRepository) Delete(otp *models.Otp) error {
	return repo.db.Delete(otp).Error
}

func (repo *OtpRepository) DeleteAllByUserId(userId string) error {
	return repo.db.Where("user_id = ?", userId).Delete(&models.Otp{}).Error
}

func (repo *OtpRepository) FindByUserId(userId string) ([]models.Otp, error) {
	var otps []models.Otp

	err := repo.db.Where("user_id = ?", userId).Find(&otps).Error

	if err != nil {
		return nil, err
	}

	return otps, nil
}

func (repo *OtpRepository) FindByUserIdAndType(userId string, otpType string) (*models.Otp, error) {
	var otp models.Otp

	err := repo.db.Where("user_id = ? AND otp_type = ?", userId, otpType).First(&otp).Error

	if err != nil {
		return nil, err
	}

	return &otp, nil
}
