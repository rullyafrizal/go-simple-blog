package seeds

import (
	"log"

	"github.com/rullyafrizal/go-simple-blog/models"
	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) error {
	categories := []string{
		"Technology",
		"Business",
		"Entertainment",
		"Fashion",
		"Sports",
		"Health",
		"Science",
		"Politics",
		"Travel",
		"Lifestyle",
		"Others",
	}

	var err error

	for _, v := range categories {
		category := models.Category{Name: v}

		err = db.Where(category).FirstOrCreate(&category).Error
	}

	log.Println("Seeded categories")

	return err
}