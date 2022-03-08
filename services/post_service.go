package services

import (
	"errors"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

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

func GetRecentPosts(c *gin.Context, limit int) ([]*models.Post, error) {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))

	posts, err := postRepository.GetRecentPosts(limit)

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

func GetMyPosts(c *gin.Context) ([]*models.Post, error) {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))

	userId, _ := utils.ExtractTokenID(c)

	posts, err := postRepository.GetPostsByUserId(uint64(userId))

	if err != nil {
		return nil, err
	}

	return posts, nil
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

	var imagePath string
	if _, err := c.FormFile("file"); err == nil {
		imagePath, err = UploadFile(c)

		if err != nil {
			return err
		}
	}

	postSlug := slug.Make(request.Title)

	if _, err := postRepository.GetPostBySlug(postSlug); err == nil {
		return errors.New("post title already exists")
	}

	post := models.Post{
		Title:      request.Title,
		Content:    template.HTML(request.Content),
		CategoryId: request.CategoryId,
		Slug:       postSlug,
		UserId:     uint64(userId),
		Image:      imagePath,
	}

	if err := postRepository.InsertPost(&post); err != nil {
		return err
	}

	return nil
}

func UpdatePost(c *gin.Context) error {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))
	var request requests.StorePostRequest
	userId, _ := utils.ExtractTokenID(c)

	if err := c.Bind(&request); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	post, err := postRepository.GetPostById(uint64(id))

	if err != nil {
		return err
	}

	if post.UserId != uint64(userId) {
		return errors.New("you are not authorized to update this post")
	}

	postSlug := slug.Make(request.Title)

	anotherPost, err := postRepository.GetPostBySlug(postSlug)

	if err != nil {
		return err
	}

	if anotherPost.ID != uint64(id) {
		return errors.New("post title already exists")
	}

	if file, _ := c.FormFile("file"); file != nil {
		imagePath, err := UploadFile(c)

		if err != nil {
			return err
		}

		post.Image = imagePath
	}

	post.Title = request.Title
	post.Content = template.HTML(request.Content)
	post.CategoryId = request.CategoryId
	post.Slug = postSlug
	post.Category = models.Category{
		ID: request.CategoryId,
	}

	if err := postRepository.UpdatePost(post); err != nil {
		return err
	}

	return nil
}

func DeletePost(c *gin.Context) error {
	postRepository := repositories.NewPostRepository(c.MustGet("db").(*gorm.DB))

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	userId, err := utils.ExtractTokenID(c)

	if err != nil {
		return err
	}

	post, err := postRepository.GetPostById(uint64(id))

	if err != nil {
		return err
	}

	if post.UserId != uint64(userId) {
		return errors.New("you are not authorized to delete this post")
	}

	if err := postRepository.DeletePost(uint64(id)); err != nil {
		return err
	}

	os.Remove("./public" + post.Image)

	return nil
}

func UploadFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")

	if err != nil {
		return "", err
	}

	extension := filepath.Ext(file.Filename)

	if extension != ".png" && extension != ".jpg" && extension != ".jpeg" {
		return "", errors.New("invalid file extension")
	}

	fileName := utils.GenerateRandomString(10) + extension
	filePath := "/images/" + fileName

	if err := c.SaveUploadedFile(file, "./public/images/"+fileName); err != nil {
		return "", err
	}

	return filePath, nil
}
