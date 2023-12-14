package services

import (
	"errors"
	"gosuper/app/helpers"
	"gosuper/app/models"
	"gosuper/app/repositories"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OtpService struct {
	otpRepository *repositories.OtpRepository
}

func NewOtpService(otpRepository *repositories.OtpRepository) *OtpService {
	return &OtpService{
		otpRepository: otpRepository,
	}
}

func (service *OtpService) GenerateOtp(userId uuid.UUID, otpType string, ttl time.Duration) (string, error) {
	existsingOtp, err := service.otpRepository.FindByUserIdAndType(userId.String(), otpType)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	if existsingOtp != nil {
		existsingOtp.Code = helpers.GenerateRandomString(6)
		existsingOtp.ValidUntil = time.Now().Add(ttl)

		err = service.otpRepository.Update(existsingOtp)

		if err != nil {
			return "", err
		}

		return existsingOtp.Code, nil
	} else {
		var otp models.Otp

		otp.ID = uuid.New()
		otp.Code = helpers.GenerateRandomString(6)
		otp.OtpType = otpType
		otp.UserID = userId
		otp.ValidUntil = time.Now().Add(ttl)

		err := service.otpRepository.Create(&otp)

		if err != nil {
			return "", err
		}

		return otp.Code, nil
	}
}

func (service *OtpService) ValidateOtp(otpType string, code string) (*models.Otp, error) {
	otp, err := service.otpRepository.FindByCodeAndType(code, otpType)

	if err != nil {
		return nil, errors.New("invalid otp")
	}

	if time.Now().After(otp.ValidUntil) {
		return nil, errors.New("otp expired")
	}

	return otp, nil
}

func (service *OtpService) DeleteOtp(otp *models.Otp) error {
	return service.otpRepository.Delete(otp)
}