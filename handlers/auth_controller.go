package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/requests"
	"github.com/rullyafrizal/go-simple-blog/services"
)

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

	c.Redirect(301, "/")

	return
}
