package repositories

import (
	"gosuper/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.Database.Create(user).Error
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := repo.Database.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) FindById(id string) (*models.User, error) {
	var user models.User

	err := repo.Database.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Update(user *models.User) error {
	return repo.Database.Save(user).Error
}

func (repo *UserRepository) Delete(user *models.User) error {
	return repo.Database.Delete(user).Error
}

func (repo *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User

	err := repo.Database.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) FindAllPaginate(page int, perPage int, search string) ([]*models.User, error) {
	var users []*models.User

	offset := (page - 1) * perPage

	query := repo.Database.Offset(offset).Limit(perPage)

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query = query.Find(&users)

	err := query.Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
