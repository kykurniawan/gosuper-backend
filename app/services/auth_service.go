package services

import (
	"gosuper/app/http/requests"
	"gosuper/app/libs/hash"
	"gosuper/app/models"

	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (service *AuthService) Login() {

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
