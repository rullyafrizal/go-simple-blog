package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/services"
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

func IndexPostsPage(c *gin.Context) {
	categories, err := services.GetAllCategories(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
		return
	}

	c.HTML(http.StatusOK, "posts/index.html", gin.H{
		"route":      "/posts",
		"title":      "Discover Posts",
		"categories": categories,
	})
}

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"route": "/",
		"title": "Go Blog",
	})
}

func ShowPostPage(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/show.html", gin.H{
		"route": "/posts/contoh-post-1",
		"title": "Contoh Post 1",
	})
}

func CreatePostPage(c *gin.Context) {
	categories, err := services.GetAllCategories(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
		return
	}

	c.HTML(http.StatusOK, "posts/create.html", gin.H{
		"route":      "/posts/create",
		"title":      "Create Post",
		"categories": categories,
	})
}

func EditPostPage(c *gin.Context) {
	categories, err := services.GetAllCategories(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
	}

	c.HTML(http.StatusOK, "posts/edit.html", gin.H{
		"route":      "/posts/edit",
		"title":      "Edit Post",
		"categories": categories,
	})
}

func DashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"route": "/dashboard",
		"title": "Dashboard",
	})
}
