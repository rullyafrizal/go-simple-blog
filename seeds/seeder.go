package seeds

import "gorm.io/gorm"

func Seed(db *gorm.DB) error {
	err := SeedCategory(db)

	return err
}