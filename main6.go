package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示如何处理静态文件
func main() {
	r := gin.Default()

	// 定义静态文件目录
	r.Static("/xxx", "./main6Templates/statics")
	// 解析模板
	r.LoadHTMLGlob("main6Templates/**/*")
	r.GET("main6/pos/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pos/index.tmpl", gin.H{
			"title": "这是一个标题",
		})
	})
	_ = r.Run(":9090")
}
