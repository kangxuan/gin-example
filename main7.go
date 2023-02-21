package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Demo
func main() {
	r := gin.Default()

	// 定义静态文件的地址
	r.Static("/xxx", "./main7Templates/statics")
	// 解析模板
	r.LoadHTMLFiles("./main7Templates/index.html")
	r.GET("/main7/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) // 渲染模板
	})

	_ = r.Run(":9090")
}
