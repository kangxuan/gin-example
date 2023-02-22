package main

import "github.com/gin-gonic/gin"

// gin第一个例子
func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//GET：请求方式；/hello：请求的路径
	r.GET("hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// 启动http服务
	_ = r.Run()
}
