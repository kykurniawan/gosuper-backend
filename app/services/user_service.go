package services

import (
	"database/sql"
	"gosuper/app/http/requests"
	"gosuper/app/models"
	"gosuper/app/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) CreateUser(
	name string,
	email string,
	password string,
) (*models.User, error) {
	user := models.User{
		ID:              uuid.New(),
		Name:            name,
		Email:           email,
		Password:        password,
		EmailVerifiedAt: sql.NullTime{},
	}

	err := service.userRepository.Create(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) GetByEmail(email string) (*models.User, error) {
	return service.userRepository.FindByEmail(email)
}

func (service *UserService) IsEmailExists(email string) bool {
	_, err := service.userRepository.FindByEmail(email)

	return err != gorm.ErrRecordNotFound
}

func (service *UserService) GetAllPaginate(request *requests.IndexRequest) ([]*models.User, error) {
	return service.userRepository.FindAllPaginate(request.Page, request.PerPage, request.Search)
}
