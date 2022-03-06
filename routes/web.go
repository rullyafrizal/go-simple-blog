package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/handlers"
	"github.com/rullyafrizal/go-simple-blog/middleware"
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
	r.GET("/posts/:slug", handlers.ShowPost)

	// Auth
	guestMiddlewaredRoute := r.Group("/auth")
	{
		guestMiddlewaredRoute.Use(middleware.GuestMiddleware())
		guestMiddlewaredRoute.GET("/login", handlers.LoginPage)
		guestMiddlewaredRoute.POST("/login", handlers.Login)
		guestMiddlewaredRoute.GET("/register", handlers.RegisterPage)
		guestMiddlewaredRoute.POST("/register", handlers.Register)
	}

	jwtMiddlewaredRoute := r.Group("")
	{
		jwtMiddlewaredRoute.Use(middleware.JwtAuthMiddleware())

		jwtMiddlewaredRoute.GET("/auth/logout", handlers.Logout)

		// Dashboard
		jwtMiddlewaredRoute.GET("/dashboard", handlers.DashboardPage)

		// Posts
		jwtMiddlewaredRoute.GET("/posts/create", handlers.CreatePostPage)
		jwtMiddlewaredRoute.POST("/posts", handlers.StorePost)
		jwtMiddlewaredRoute.GET("/posts/edit/1", handlers.EditPostPage)

		jwtMiddlewaredRoute.GET("/auth/me", handlers.GetAuthenticatedUser)
	}

	port := utils.Getenv("PORT", "8080")

	r.Run(":" + port)
}
