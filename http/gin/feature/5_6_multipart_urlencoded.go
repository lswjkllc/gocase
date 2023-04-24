package feature

import "github.com/gin-gonic/gin"

func MultipartAndUrlencodedForm(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func MultipartAndUrlencodedBinding(c *gin.Context) {
	// 你可以使用显式绑定声明绑定 multipart form：
	// c.ShouldBindWith(&form, binding.Form)
	// 或者简单地使用 ShouldBind 方法自动绑定：
	var form LoginForm
	// 在这种情况下，将自动选择合适的绑定
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
