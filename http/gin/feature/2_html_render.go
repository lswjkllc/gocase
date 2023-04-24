package feature

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "Posts",
	})
}

func UsersIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "Users",
	})
}
