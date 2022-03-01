package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/utils"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.LoadHTMLGlob("views/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	port := utils.Getenv("PORT", "8080")

	r.Run(":" + port)
}
