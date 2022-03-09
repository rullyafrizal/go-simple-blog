package repositories

import (
	"github.com/rullyafrizal/go-simple-blog/models"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{
		DB: db,
	}
}

func (repository *PostRepositoryImpl) GetAllPosts(query string) ([]*models.Post, error) {
	var posts []*models.Post

	// TODO: add search query logic
	q := repository.DB.Preload("Category").Preload("User")

	if query != "" {
		q = q.Where("title LIKE ?", "%"+query+"%")
	}

	err := q.Order("created_at desc").Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (repository *PostRepositoryImpl) GetRecentPosts(limit int) ([]*models.Post, error) {
	var posts []*models.Post

	err := repository.DB.Preload("Category").Preload("User").Limit(limit).Order("created_at desc").Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (repository *PostRepositoryImpl) GetPostById(id uint64) (*models.Post, error) {
	var post *models.Post

	err := repository.DB.Where("id = ?", id).Preload("Category").Preload("User").First(&post).Error

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (repository *PostRepositoryImpl) GetPostsByUserId(userId uint64) ([]*models.Post, error) {
	var posts []*models.Post

	err := repository.DB.Where("user_id = ?", userId).Preload("Category").Preload("User").Order("created_at desc").Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (repository *PostRepositoryImpl) GetPostBySlug(slug string) (*models.Post, error) {
	var post *models.Post

	err := repository.DB.Where("slug = ?", slug).Preload("Category").Preload("User").First(&post).Error

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (repository *PostRepositoryImpl) InsertPost(post *models.Post) error {
	return repository.DB.Create(post).Error
}

func (repository *PostRepositoryImpl) UpdatePost(post *models.Post) error {
	return repository.DB.Save(post).Error
}

func (repository *PostRepositoryImpl) DeletePost(id uint64) error {
	return repository.DB.Delete(&models.Post{}, id).Error
}
