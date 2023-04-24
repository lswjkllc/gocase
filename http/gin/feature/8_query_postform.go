package feature

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func QueryAndPostForm(c *gin.Context) {

	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
