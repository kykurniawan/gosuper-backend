package services

import (
	"gosuper/app/helpers"
	"gosuper/app/http/requests"
	"gosuper/app/libs/hash"
	"gosuper/app/models"
	"gosuper/app/repositories"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	userService            *UserService
	refreshTokenRepository *repositories.RefreshTokenRepository
}

func NewAuthService(userService *UserService, refreshTokenRepository *repositories.RefreshTokenRepository) *AuthService {
	return &AuthService{
		userService:            userService,
		refreshTokenRepository: refreshTokenRepository,
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
	refreshToken := models.RefreshToken{
		ID:         uuid.New(),
		UserID:     user.ID,
		Token:      helpers.GenerateRandomString(32),
		ValidUntil: time.Now().Add(time.Hour * 24 * 7),
	}

	err := service.refreshTokenRepository.Create(&refreshToken)

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (service *AuthService) GenerateAccessToken(user *models.User) (string, error) {
	expiresMilisecond, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRES"))

	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Microsecond * time.Duration(expiresMilisecond)).Unix(),
		"iss": "gosuper",
		"aud": "gosuper",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
