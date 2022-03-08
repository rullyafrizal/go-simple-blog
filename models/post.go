package models

import (
	"database/sql"
	"html/template"

	"github.com/microcosm-cc/bluemonday"
)

type Post struct {
	ID          uint64        `json:"id" gorm:"primaryKey"`
	Title       string        `json:"title" gorm:"size:255;not null"`
	Content     template.HTML `json:"content" gorm:"not null"`
	Image       string        `json:"image" gorm:"size:255"`
	Slug        string        `json:"slug" gorm:"size:255;uniqueIndex"`
	UserId      uint64        `json:"user_id"`
	CategoryId  uint64        `json:"category_id"`
	Category    Category      `json:"category"`
	User        User          `json:"user"`
	CreatedAt   sql.NullTime  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   sql.NullTime  `json:"updated_at" gorm:"autoUpdateTime"`
}

func (p *Post) SanitizeContent() {
	sanitizer := bluemonday.UGCPolicy()

	p.Content = template.HTML(sanitizer.Sanitize(string(p.Content)))
}
