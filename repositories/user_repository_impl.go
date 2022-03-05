package repositories

import (
	"github.com/rullyafrizal/go-simple-blog/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) CreateUser(user *models.User) error {
	return repository.DB.Create(user).Error
}

func (repository *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := repository.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) GetUserById(id uint64) (*models.User, error) {
	var user models.User

	err := repository.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	err := repository.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}