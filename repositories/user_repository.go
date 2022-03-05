package repositories

import "github.com/rullyafrizal/go-simple-blog/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint64) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}