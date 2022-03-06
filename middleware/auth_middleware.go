package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.TokenValid(c); err != nil {
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
			return
		}

		c.Next()
	}
}

func GuestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.TokenValid(c); err == nil {
			c.Redirect(http.StatusFound, "/")

			return
		}

		c.Next()
	}
}
