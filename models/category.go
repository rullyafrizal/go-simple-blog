package models

type Category struct {
	ID uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Posts []Post `json:"posts"`
}