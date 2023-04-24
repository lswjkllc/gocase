package feature

import "github.com/gin-gonic/gin"

/*
通常，JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c。
如果要按字面对这些字符进行编码，则可以使用 PureJSON。Go 1.6 及更低版本无法使用此功能。
*/

func UnicodeEntity(c *gin.Context) {
	c.JSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

func LeterEntity(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}
