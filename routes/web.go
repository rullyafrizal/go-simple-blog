package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/handlers"
	"github.com/rullyafrizal/go-simple-blog/utils"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.Static("/assets", "./resources/assets")
	r.Static("/vendor", "./resources/vendor")
	r.LoadHTMLGlob("resources/templates/**/*")

	// Homepage
	r.GET("/", handlers.IndexPage)

	// Posts
	r.GET("/posts", handlers.IndexPostsPage)
	r.GET("/posts/contoh-post-1", handlers.ShowPostPage)
	r.GET("/posts/create", handlers.CreatePostPage)
	r.GET("/posts/edit/1", handlers.EditPostPage)

	// Auth
	r.GET("/auth/login", handlers.LoginPage)
	r.GET("/auth/register", handlers.RegisterPage)
	r.POST("/auth/register", handlers.Register)
	
	// Dashboard
	r.GET("/dashboard", handlers.DashboardPage)

	port := utils.Getenv("PORT", "8080")

	r.Run(":" + port)
}
