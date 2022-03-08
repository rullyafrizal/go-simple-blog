package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/requests"
	"github.com/rullyafrizal/go-simple-blog/services"
	"github.com/rullyafrizal/go-simple-blog/utils"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.html", gin.H{
		"route": "auth/login",
		"title": "Login",
	})
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{
		"route": "auth/register",
		"title": "Register",
	})
}

func Login(c *gin.Context) {
	var loginRequest requests.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
	}

	errors := loginRequest.Validate()

	if len(errors) != 0 {
		c.HTML(http.StatusUnprocessableEntity, "auth/login.html", gin.H{
			"route":  "auth/login",
			"title":  "Login",
			"errors": errors,
		})

		return
	}

	if err := services.Login(c, &loginRequest); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "auth/login.html", gin.H{
			"route": "auth/login",
			"title": "Login",
			"errors": map[string]string{
				"error": err.Error(),
			},
		})

		return
	}

	c.Redirect(http.StatusFound, "/")
}

func Register(c *gin.Context) {
	var registerRequest requests.StoreUserRequest

	if err := c.Bind(&registerRequest); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
	}

	errors := registerRequest.Validate()

	if len(errors) != 0 {
		c.HTML(http.StatusUnprocessableEntity, "auth/register.html", gin.H{
			"route":  "auth/register",
			"title":  "Register",
			"errors": errors,
		})

		return
	}

	if err := services.Register(c, &registerRequest); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "auth/register.html", gin.H{
			"route": "auth/register",
			"title": "Register",
			"errors": map[string]string{
				"error": err.Error(),
			},
		})

		return
	}

	// Login after successfully registered
	if err := services.Login(c, &requests.LoginRequest{Email: registerRequest.Email, Password: registerRequest.Password}); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "auth/register.html", gin.H{
			"route": "auth/register",
			"title": "Register",
			"errors": map[string]string{
				"error": err.Error(),
			},
		})
	}

	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {
	utils.RemoveCookie(c, "blog_jwt_token")

	c.Redirect(http.StatusFound, "/auth/login")
}
