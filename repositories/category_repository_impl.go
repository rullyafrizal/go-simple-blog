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

	err := repository.DB.Where("id = ?", id).First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (repository *CategoryRepositoryImpl) InsertCategory(category *models.Category) error {
	return repository.DB.Create(category).Error
}

func (repository *CategoryRepositoryImpl) UpdateCategory(category *models.Category) error {
	return repository.DB.Save(category).Error
}

func (repository *CategoryRepositoryImpl) DeleteCategory(id uint64) error {
	return repository.DB.Delete(&models.Category{}, id).Error
}

