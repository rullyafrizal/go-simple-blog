package repositories

import "github.com/rullyafrizal/go-simple-blog/models"

type CategoryRepository interface {
	GetAllCategories() ([]*models.Category, error)
	GetCategoryById(id uint64) (*models.Category, error)
	InsertCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint64) error
}