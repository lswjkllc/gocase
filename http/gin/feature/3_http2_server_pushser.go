package feature

import (
	"log"

	"github.com/gin-gonic/gin"
)

/*
http.Pusher 仅支持 go1.8+。 更多信息，请查阅 golang blog: https://go.dev/blog/h2push。
*/

func Http2ServerPusher(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		// 使用 pusher.Push() 做服务器推送
		if err := pusher.Push("/assets/app.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}

	c.HTML(200, "https", gin.H{
		"status": "success",
	})
}
