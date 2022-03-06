package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/models"
	"github.com/rullyafrizal/go-simple-blog/repositories"
	"github.com/rullyafrizal/go-simple-blog/requests"
	"github.com/rullyafrizal/go-simple-blog/utils"
	"gorm.io/gorm"
)

func Register(c *gin.Context, request *requests.StoreUserRequest) error {
	userRepository := repositories.NewUserRepository(c.MustGet("db").(*gorm.DB))

	email := request.Email
	name := request.Name
	password := request.Password

	_, err := userRepository.GetUserByEmail(email)

	if err == nil {
		return errors.New("email already exists")
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err = userRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func Login(c *gin.Context, request *requests.LoginRequest) error {
	userRepository := repositories.NewUserRepository(c.MustGet("db").(*gorm.DB))
	user, err := userRepository.GetUserByEmail(request.Email)

	if err != nil {
		return errors.New("email not found")
	}

	if !user.CheckPasswordHash(request.Password) {
		return errors.New("invalid Password")
	}

	token, err := utils.GenerateToken(int64(user.ID))

	if err != nil {
		return err
	}

	utils.SetCookie(c, "blog_jwt_token", token)

	return nil
}

func GetAuthenticatedUser(c *gin.Context) (*models.User, error) {
	userRepository := repositories.NewUserRepository(c.MustGet("db").(*gorm.DB))
	id, err := utils.ExtractTokenID(c)

	if err != nil {
		return nil, err
	}

	user, err := userRepository.GetUserById(uint64(id))

	if err != nil {
		return nil, err
	}

	return user, nil
}
