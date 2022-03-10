package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/go-simple-blog/requests"
	"github.com/rullyafrizal/go-simple-blog/services"
	"github.com/rullyafrizal/go-simple-blog/utils"
)

func IndexPostsPage(c *gin.Context) {
	categories, err := services.GetAllCategories(c)
	user, err := services.GetAuthenticatedUser(c)
	posts, err := services.GetAllPosts(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
		return
	}

	recentPosts, _ := services.GetRecentPosts(c, 3)

	c.HTML(http.StatusOK, "posts/index.html", gin.H{
		"route":        "/posts",
		"title":        "Discover Posts",
		"posts":        posts,
		"categories":   categories,
		"recent_posts": recentPosts,
		"user":         user,
	})
}

func ShowPost(c *gin.Context) {
	categories, err := services.GetAllCategories(c)
	post, err := services.GetPostBySlug(c)

	post.SanitizeContent()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
		return
	}

	user, _ := services.GetAuthenticatedUser(c)
	recentPosts, _ := services.GetRecentPosts(c, 3)

	c.HTML(http.StatusOK, "posts/show.html", gin.H{
		"route":        "/posts/" + post.Slug,
		"title":        post.Title,
		"user":         user,
		"post":         post,
		"recent_posts": recentPosts,
		"categories":   categories,
	})
}

func GetPostsByCategory(c *gin.Context) {
	category, err := services.GetCategoryWithPosts(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
		return
	}

	user, _ := services.GetAuthenticatedUser(c)
	recentPosts, _ := services.GetRecentPosts(c, 3)
	categories, _ := services.GetAllCategories(c)

	c.HTML(http.StatusOK, "category-posts.html", gin.H{
		"route":        "/categories/" + strconv.Itoa(int(category.ID)) + "/posts",
		"title":        category.Name,
		"user":         user,
		"categories":   categories,
		"recent_posts": recentPosts,
		"category":     category,
	})

}

func StorePost(c *gin.Context) {
	user, _ := services.GetAuthenticatedUser(c)
	categories, _ := services.GetAllCategories(c)
	var request requests.StorePostRequest

	if err := c.Bind(&request); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
		return
	}

	if errors := request.Validate(c); len(errors) > 0 {
		c.HTML(http.StatusUnprocessableEntity, "posts/create.html", gin.H{
			"route":      "posts/create",
			"title":      "Create Post",
			"categories": categories,
			"user":       user,
			"errors":     errors,
		})
		return
	}

	err := services.CreatePost(c, &request)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})

		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}

func DestroyPost(c *gin.Context) {
	err := services.DeletePost(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}

func CreatePostPage(c *gin.Context) {
	user, _ := services.GetAuthenticatedUser(c)
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
		"user":       user,
	})
}

func EditPostPage(c *gin.Context) {
	user, _ := services.GetAuthenticatedUser(c)
	categories, err := services.GetAllCategories(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
	}

	post, _ := services.GetPostById(c)
	userId, err := utils.ExtractTokenID(c)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
	}

	if uint64(userId) != post.UserId {
		c.Redirect(http.StatusFound, "/dashboard")
	}

	c.HTML(http.StatusOK, "posts/edit.html", gin.H{
		"route":      "/posts/edit",
		"title":      "Edit Post",
		"categories": categories,
		"user":       user,
		"post":       post,
	})
}

func UpdatePost(c *gin.Context) {
	var request requests.UpdatePostRequest

	if err := c.Bind(&request); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "Internal Server Error",
		})
		return
	}

	postId, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if errors := request.Validate(c, postId); len(errors) > 0 {
		user, _ := services.GetAuthenticatedUser(c)
		categories, _ := services.GetAllCategories(c)
		post, _ := services.GetPostById(c)

		c.HTML(http.StatusUnprocessableEntity, "posts/edit.html", gin.H{
			"route":      "/posts/edit",
			"title":      "Edit Post",
			"categories": categories,
			"user":       user,
			"post":       post,
			"errors":     errors,
		})
		return
	}

	err := services.UpdatePost(c, &request)

	if err != nil {
		user, _ := services.GetAuthenticatedUser(c)
		categories, _ := services.GetAllCategories(c)
		post, _ := services.GetPostById(c)

		c.HTML(http.StatusUnprocessableEntity, "posts/edit.html", gin.H{
			"route":      "posts/edit",
			"title":      "Edit Post",
			"categories": categories,
			"user":       user,
			"post":       post,
			"errors": map[string]string{
				"error": err.Error(),
			},
		})

		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}
