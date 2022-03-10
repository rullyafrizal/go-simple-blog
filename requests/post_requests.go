package requests

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/rullyafrizal/go-simple-blog/repositories"
	"gorm.io/gorm"
)

type StorePostRequest struct {
	Title      string `form:"title"`
	Content    string `form:"content"`
	CategoryId uint64 `form:"category_id"`
}

type UpdatePostRequest struct {
	Title      string `form:"title"`
	Content    string `form:"content"`
	CategoryId uint64 `form:"category_id"`
}

func (r *StorePostRequest) Validate(c *gin.Context) map[string]string {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))
	var errors map[string]string = make(map[string]string)

	if strings.TrimSpace(r.Title) == "" {
		errors["title"] = "Title can't be blank"
	} else {
		postSlug := slug.Make(r.Title)

		if _, err := postRepository.GetPostBySlug(postSlug); err == nil {
			errors["title"] = "Title already exists"
		}
	}

	if strings.TrimSpace(r.Content) == "" {
		errors["content"] = "Content can't be blank"
	}

	if r.CategoryId == 0 {
		errors["category_id"] = "Category can't be blank"
	}

	return errors
}

func (r *UpdatePostRequest) Validate(c *gin.Context, postId uint64) map[string]string {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))
	var errors map[string]string = make(map[string]string)

	if strings.TrimSpace(r.Title) == "" {
		errors["title"] = "Title can't be blank"
	} else {
		postSlug := slug.Make(r.Title)

		anotherPost, _ := postRepository.GetPostBySlug(postSlug)

		if anotherPost != nil && anotherPost.ID != postId {
			errors["title"] = "Title already exists"
		}
	}

	if strings.TrimSpace(r.Content) == "" {
		errors["content"] = "Content can't be blank"
	}

	if r.CategoryId == 0 {
		errors["category_id"] = "Category can't be blank"
	}

	return errors
}
