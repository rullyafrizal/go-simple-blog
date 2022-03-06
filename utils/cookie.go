package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetCookie(ctx *gin.Context, name string, value string) {
	maxAge, err := strconv.Atoi(Getenv("COOKIE_MAX_AGE", "3600"))

	if err != nil {
		maxAge = 3600
	}

	ctx.SetCookie(name, value, maxAge, "/", "", false, true)
}

func GetCookie(ctx *gin.Context, name string) string {
	val, err := ctx.Cookie(name)

	if err != nil {
		return ""
	}

	return val
}

func RemoveCookie(ctx *gin.Context, name string) {
	ctx.SetCookie(name, "", -1, "/", "", false, true)
}
