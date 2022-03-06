package models

import (
	"database/sql"
)

type Post struct {
	ID          uint64       `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title" gorm:"size:255"`
	Content     string       `json:"content"`
	Image       string       `json:"image" gorm:"size:255"`
	Slug        string       `json:"slug" gorm:"size:255;uniqueIndex"`
	UserId      uint64       `json:"user_id"`
	CategoryId  uint64       `json:"category_id"`
	PublishedAt sql.NullTime `json:"published_at"`
	CreatedAt   sql.NullTime `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`
}
