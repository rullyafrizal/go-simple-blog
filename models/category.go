package models

type Category struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"size:255;not null"`
	Posts []Post `json:"posts"`
}
