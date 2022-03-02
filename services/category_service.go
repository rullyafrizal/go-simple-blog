package services

import (
	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/models"
	"github.com/rullyafrizal/go-simple-blog/repositories"
	"gorm.io/gorm"
)

func GetAllCategories(ctx *gin.Context) ([]*models.Category, error) {
	categoryRepository := repositories.NewCategoryRepository(ctx.MustGet("db").(*gorm.DB))

	categories, err := categoryRepository.GetAllCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}