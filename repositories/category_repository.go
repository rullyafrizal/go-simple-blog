package repositories

import "github.com/rullyafrizal/go-simple-blog/models"

type CategoryRepository interface {
	GetAllCategories() ([]*models.Category, error)
	GetCategoryById(id uint64) (*models.Category, error)
}