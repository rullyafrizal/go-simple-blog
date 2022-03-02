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

	r.GET("/", handlers.Index)
	r.GET("/posts", handlers.IndexPosts)
	r.GET("/auth/login", handlers.LoginPage)
	r.GET("/auth/register", handlers.RegisterPage)
	r.GET("/posts/contoh-post-1", handlers.ShowPost)

	port := utils.Getenv("PORT", "8080")

	r.Run(":" + port)
}
