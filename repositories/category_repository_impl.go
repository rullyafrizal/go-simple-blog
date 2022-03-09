package repositories

import (
	"github.com/rullyafrizal/go-simple-blog/models"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (repository *CategoryRepositoryImpl) GetAllCategories() ([]*models.Category, error) {
	var categories []*models.Category

	err := repository.DB.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (repository *CategoryRepositoryImpl) GetCategoryById(id uint64) (*models.Category, error) {
	var category models.Category

	err := repository.DB.Where("id = ?", id).Preload("Posts.User").First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}
