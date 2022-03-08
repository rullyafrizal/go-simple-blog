package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/services"
)

func HomePage(c *gin.Context) {
	user, _ := services.GetAuthenticatedUser(c)
	recentPosts, _ := services.GetRecentPosts(c, 3)
	categories, _ := services.GetAllCategories(c)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"route":        "/",
		"title":        "Go Blog",
		"user":         user,
		"posts":        recentPosts,
		"recent_posts": recentPosts,
		"categories":   categories,
	})
}

func DashboardPage(c *gin.Context) {
	user, _ := services.GetAuthenticatedUser(c)
	myPosts, _ := services.GetMyPosts(c)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"route": "/dashboard",
		"title": "Dashboard",
		"user":  user,
		"posts": myPosts,
	})
}

func GetAuthenticatedUser(c *gin.Context) {
	user, err := services.GetAuthenticatedUser(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	c.JSON(http.StatusOK, user)
}
