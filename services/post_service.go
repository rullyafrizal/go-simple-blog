package services

import (
	"errors"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/rullyafrizal/go-simple-blog/models"
	"github.com/rullyafrizal/go-simple-blog/repositories"
	"github.com/rullyafrizal/go-simple-blog/requests"
	"github.com/rullyafrizal/go-simple-blog/utils"
	"gorm.io/gorm"
)

func GetAllPosts(c *gin.Context) ([]*models.Post, error) {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))

	posts, err := postRepository.GetAllPosts()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func GetPostById(c *gin.Context) (*models.Post, error) {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return nil, err
	}

	post, err := postRepository.GetPostById(uint64(id))

	if err != nil {
		return nil, err
	}

	return post, nil
}

func GetPostBySlug(c *gin.Context) (*models.Post, error) {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))

	slug := c.Param("slug")

	post, err := postRepository.GetPostBySlug(slug)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func CreatePost(c *gin.Context) error {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))
	var request requests.StorePostRequest
	userId, _ := utils.ExtractTokenID(c)

	if err := c.Bind(&request); err != nil {
		return err
	}

	post := models.Post{
		Title:      request.Title,
		Content:    request.Content,
		CategoryId: request.CategoryId,
		Slug:       slug.Make(request.Title),
		UserId:     uint64(userId),
	}

	if request.IsPublished != "" {
		post.PublishedAt.Time = time.Now()
		post.PublishedAt.Valid = true
	}

	if err := postRepository.InsertPost(&post); err != nil {
		return err
	}

	return nil
}

func UploadFile(c *gin.Context) error {
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

	extension := filepath.Ext(file.Filename)

	if extension != ".png" && extension != ".jpg" && extension != ".jpeg" {
		return errors.New("invalid file extension")
	}

	fileName := utils.GenerateRandomString(10) + extension

	if err := c.SaveUploadedFile(file, "./public/images/"+fileName); err != nil {
		return err
	}

	return nil
}
