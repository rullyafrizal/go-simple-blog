package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/models"
	"github.com/rullyafrizal/go-simple-blog/repositories"
	"github.com/rullyafrizal/go-simple-blog/requests"
	"gorm.io/gorm"
)

func Register(c *gin.Context, request *requests.StoreUserRequest) error {
	userRepository := repositories.NewUserRepository(c.MustGet("db").(*gorm.DB))

	email := request.Email
	name := request.Name
	password := request.Password

	_, err := userRepository.GetUserByEmail(email)

	if err == nil {
		return errors.New("Email already exists")
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
		return errors.New("Email not found")
	}

	if !user.CheckPasswordHash(request.Password) {
		return errors.New("Invalid Password")
	}

	// TODO: buat mekanisme login
	// 1. Simpan JWT ke dalam cookie based session
	// 2. Cari tahu cara untuk mengaksesnya di kemudian

	// jwt, err := utils.GenerateToken(int64(user.ID))

	// sessions.S

	return nil
}
