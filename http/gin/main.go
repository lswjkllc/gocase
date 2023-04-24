package main

import (
	"context"
	"gocase/http/gin/feature"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// // HTTP2 Server Pusher
// var html = template.Must(template.New("https").Parse(`
// <html>
// <head>
//   <title>Https Test</title>
//   <script src="/assets/app.js"></script>
// </head>
// <body>
//   <h1 style="color:red;">Welcome, Ginner!</h1>
// </body>
// </html>
// `))

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	router := gin.Default()

	// AsciiJSON
	router.GET("/ascii/json", feature.AsciiJSON)

	// HTML render
	router.LoadHTMLGlob("./http/gin/templates/**/*")
	router.GET("/posts/index", feature.PostsIndex)
	router.GET("/users/index", feature.UsersIndex)

	// // HTTP2 Server Pusher
	// router.Static("/assets", "./assets")
	// router.SetHTMLTemplate(html)
	// router.GET("/", feature.Http2ServerPusher)
	// // 监听并在 https://127.0.0.1:8080 上启动服务
	// router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")

	// JSONP
	router.GET("/jsonp", feature.Jsonp)

	// Multipart/Urlencoded
	router.POST("/login", feature.MultipartAndUrlencodedBinding)
	router.POST("/form_post", feature.MultipartAndUrlencodedForm)

	// PureJSON
	// 提供 unicode 实体
	router.GET("/json", feature.UnicodeEntity)
	// 提供字面字符
	router.GET("/purejson", feature.LeterEntity)

	// Query and Post Form
	router.POST("/post", feature.QueryAndPostForm)

	// 你也可以使用自己的 SecureJSON 前缀
	// router.SecureJsonPrefix(")]}',\n")
	router.GET("/securejson", feature.SecureJson)

	/* 自定义验证器 Custom Validator */
	// 注册
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", feature.BookableDateValidator)
	}
	// 验证
	router.GET("/bookable", feature.GetBookable)

	// // 监听并在 0.0.0.0:8080 上启动服务
	// router.Run(":8080")

	/* Graceful Restart or Stop */
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// 启动 Goroutine 监听服务
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	// 最大时间控制，通知该服务端它有 5 秒的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
