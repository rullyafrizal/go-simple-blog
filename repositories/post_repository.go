package repositories

import "github.com/rullyafrizal/go-simple-blog/models"

type PostRepository interface {
	GetAllPosts() ([]*models.Post, error)
	GetRecentPosts(limit int) ([]*models.Post, error)
	GetPostById(id uint64) (*models.Post, error)
	GetPostsByUserId(userId uint64) ([]*models.Post, error)
	GetPostBySlug(slug string) (*models.Post, error)
	InsertPost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(id uint64) error
}