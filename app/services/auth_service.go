package services

import (
	"database/sql"
	"errors"
	"gosuper/app/constants"
	"gosuper/app/helpers"
	"gosuper/app/http/requests"
	"gosuper/app/libs/hash"
	"gosuper/app/libs/queue"
	"gosuper/app/models"
	"gosuper/app/repositories"
	"gosuper/config"
	"gosuper/resources"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	userService            *UserService
	otpService             *OtpService
	mailService            *MailService
	refreshTokenRepository *repositories.RefreshTokenRepository
	queue                  *queue.Queue
}

func NewAuthService(
	userService *UserService,
	otpService *OtpService,
	mailService *MailService,
	refreshTokenRepository *repositories.RefreshTokenRepository,
	queue *queue.Queue,
) *AuthService {
	return &AuthService{
		userService:            userService,
		otpService:             otpService,
		mailService:            mailService,
		refreshTokenRepository: refreshTokenRepository,
		queue:                  queue,
	}
}

func (service *AuthService) Login(request *requests.LoginRequest) (string, string, error) {
	user, err := service.userService.GetByEmail(request.Email)

	if err != nil {
		return "", "", fiber.NewError(fiber.StatusBadRequest, "Email or password is incorrect!")
	}

	if !hash.Compare(request.Password, user.Password) {
		return "", "", fiber.NewError(fiber.StatusBadRequest, "Email or password is incorrect!")
	}

	refreshToken, err := service.GenerateRefreshToken(user)

	if err != nil {
		return "", "", err
	}

	accessToken, err := service.GenerateAccessToken(user)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken.Token, nil

}

func (service *AuthService) RefreshToken(request *requests.RefreshTokenRequest) (string, string, error) {
	refreshToken, err := service.refreshTokenRepository.FindByToken(request.RefreshToken)

	if err != nil {
		return "", "", fiber.NewError(fiber.StatusBadRequest, "Refresh token is invalid!")
	}

	if refreshToken.ValidUntil.Before(time.Now()) {
		return "", "", fiber.NewError(fiber.StatusBadRequest, "Refresh token is expired!")
	}

	user, err := service.userService.GetById(refreshToken.UserID)

	if err != nil {
		return "", "", err
	}

	newAccessToken, err := service.GenerateAccessToken(user)

	if err != nil {
		return "", "", err
	}

	service.refreshTokenRepository.Delete(refreshToken)

	newRefreshToken, err := service.GenerateRefreshToken(user)

	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken.Token, nil
}

func (service *AuthService) Logout(request *requests.LogoutRequest) error {
	refreshToken, err := service.refreshTokenRepository.FindByToken(request.RefreshToken)

	if err != nil {
		return err
	}

	err = service.refreshTokenRepository.DeleteAllByUserId(refreshToken.UserID.String())

	if err != nil {
		return err
	}

	return nil
}

func (service *AuthService) Register(request *requests.RegisterRequest) (*models.User, error) {
	if service.userService.IsEmailExists(request.Email) {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Email already exists!")
	}
	hashedPassword, err := hash.Make(request.Password)

	if err != nil {
		return nil, err
	}

	user, err := service.userService.CreateUser(
		request.Name,
		request.Email,
		hashedPassword,
	)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (service *AuthService) GenerateRefreshToken(user *models.User) (*models.RefreshToken, error) {
	expiresSecond := config.Token.RefeshTokenExpiresIn

	userRefreshTokens, err := service.refreshTokenRepository.FindByUserId(user.ID.String())

	if err != nil {
		return nil, err
	}

	if len(userRefreshTokens) > 0 {
		for _, refreshToken := range userRefreshTokens {
			service.refreshTokenRepository.Delete(&refreshToken)
		}
	}

	refreshToken := models.RefreshToken{
		ID:         uuid.New(),
		UserID:     user.ID,
		Token:      helpers.GenerateRandomString(32),
		ValidUntil: time.Now().Add(time.Second * time.Duration(expiresSecond)),
	}

	err = service.refreshTokenRepository.Create(&refreshToken)

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (service *AuthService) GenerateAccessToken(user *models.User) (string, error) {
	expiresSecond := config.Token.AccessTokenExpiresIn

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(expiresSecond)).Unix(),
		"iss": "gosuper",
		"aud": "gosuper",
	})

	tokenString, err := token.SignedString([]byte(config.App.Key))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *AuthService) ValidateAccessToken(token string) (*models.User, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check signing method
			return nil, errors.New("invalid token")
		}

		return []byte(config.App.Key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		userID, err := uuid.Parse(claims["sub"].(string))

		if err != nil {
			return nil, err
		}

		user, err := service.userService.GetById(userID)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, errors.New("invalid token")
}

func (service *AuthService) ForgotPassword(request *requests.ForgotPasswordRequest) error {
	user, err := service.userService.GetByEmail(request.Email)

	if err != nil {
		return errors.New("email not found")
	}

	otp, err := service.otpService.GenerateOtp(user.ID, constants.ResetPasswordOtp, time.Minute*60)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	queue := queue.SendEmailQueue{
		Email:        user.Email,
		Subject:      "Reset Password OTP",
		MailTemplate: resources.ResetPasswordOtpTemplate,
		Data: struct {
			Name    string
			Otp     string
			Subject string
		}{
			Name:    user.Name,
			Otp:     otp,
			Subject: "Reset Password OTP",
		},
	}

	err = service.queue.Publish(config.Queue.Mail.RoutingKey, queue)

	if err != nil {
		return err
	}

	return nil
}

func (service *AuthService) ResetPassword(request *requests.ResetPasswordRequest) error {
	otp, err := service.otpService.ValidateOtp(constants.ResetPasswordOtp, request.Otp)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid OTP")
	}

	user, err := service.userService.GetById(otp.UserID)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "User not found")
	}

	hashedPassword, err := hash.Make(request.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = service.userService.UpdateUser(user)

	if err != nil {
		return err
	}

	err = service.otpService.DeleteOtp(otp)

	if err != nil {
		return err
	}

	queue := queue.SendEmailQueue{
		Email:        user.Email,
		Subject:      "Reset Password Success",
		MailTemplate: resources.PasswordChangedNotificationTemplate,
		Data: struct {
			Name    string
			Subject string
			Time    string
			Email   string
		}{
			Name:    user.Name,
			Subject: "Reset Password Success",
			Time:    time.Now().Format("02 Jan 2006 15:04:05"),
			Email:   config.Mail.ReplyTo,
		},
	}

	err = service.queue.Publish(config.Queue.Mail.RoutingKey, queue)

	if err != nil {
		return err
	}

	return nil
}

func (service *AuthService) ResendEmailVerification(user *models.User) error {
	otp, err := service.otpService.GenerateOtp(user.ID, constants.EmailVerificationOtp, time.Minute*60)

	if err != nil {
		return err
	}

	queue := queue.SendEmailQueue{
		Email:        user.Email,
		Subject:      "Email Verification OTP",
		MailTemplate: resources.EmailVerificationOtpTemplate,
		Data: struct {
			Name    string
			Otp     string
			Subject string
		}{
			Name:    user.Name,
			Otp:     otp,
			Subject: "Email Verification OTP",
		},
	}

	err = service.queue.Publish(config.Queue.Mail.RoutingKey, queue)

	if err != nil {
		return err
	}

	return nil
}

func (service *AuthService) VerifyEmail(request *requests.VerifyEmailRequest) (*models.User, error) {
	otp, err := service.otpService.ValidateOtp(constants.EmailVerificationOtp, request.Otp)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid OTP")
	}

	user, err := service.userService.GetById(otp.UserID)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "User not found")
	}

	user.EmailVerifiedAt = sql.NullTime{
		Time: time.Now(),
		Valid: true,
	}

	err = service.userService.UpdateUser(user)

	if err != nil {
		return nil, err
	}

	err = service.otpService.DeleteOtp(otp)

	if err != nil {
		return nil, err
	}

	return user, nil
}
