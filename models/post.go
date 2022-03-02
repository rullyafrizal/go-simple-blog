package models

import "database/sql"

type Post struct {
	ID          uint64       `json:"id"`
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	Image       string       `json:"image"`
	Slug        string       `json:"slug"`
	UserId      uint64       `json:"user_id"`
	CategoryId  uint64       `json:"category_id"`
	PublishedAt sql.NullTime `json:"published_at"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}
