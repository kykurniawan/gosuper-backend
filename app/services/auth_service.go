package services

import (
	"errors"
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
	expiresSecond, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRES"))

	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(expiresSecond)).Unix(),
		"iss": "gosuper",
		"aud": "gosuper",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

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

		return []byte(os.Getenv("JWT_SECRET")), nil
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
