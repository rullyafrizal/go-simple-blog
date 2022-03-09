package services

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/models"
	"github.com/rullyafrizal/go-simple-blog/repositories"
	"gorm.io/gorm"
)

func GetAllCategories(c *gin.Context) ([]*models.Category, error) {
	categoryRepository := repositories.NewCategoryRepository(c.MustGet("db").(*gorm.DB))

	categories, err := categoryRepository.GetAllCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryWithPosts(c *gin.Context) (*models.Category, error) {
	categoryRepository := repositories.NewCategoryRepository(c.MustGet("db").(*gorm.DB))

	id, _ := strconv.Atoi(c.Param("id"))

	category, err := categoryRepository.GetCategoryById(uint64(id))

	if err != nil {
		return nil, err
	}

	return category, nil
}
