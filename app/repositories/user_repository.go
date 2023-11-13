package repositories

import (
	"gosuper/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := repo.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) FindById(id string) (*models.User, error) {
	var user models.User

	err := repo.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Update(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) Delete(user *models.User) error {
	return repo.db.Delete(user).Error
}

func (repo *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User

	err := repo.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) FindAllPaginate(page int, perPage int, search string) ([]*models.User, error) {
	var users []*models.User

	offset := (page - 1) * perPage

	query := repo.db.Offset(offset).Limit(perPage)

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
