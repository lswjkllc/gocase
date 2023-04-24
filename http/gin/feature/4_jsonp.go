package feature

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
使用 JSONP 向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。
*/

func Jsonp(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}

	// /JSONP?callback=x
	// 将输出：x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)
}
