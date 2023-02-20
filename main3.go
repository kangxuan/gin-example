package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 实例化gin框架
	r := gin.Default()

	// 定义模板
	// 解析模板
	r.LoadHTMLFiles("main3Templates/index.tmpl")
	r.GET("/main3/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{ // // 渲染模板
			"title": "这是一个标题",
		})
	})

	// 运行
	_ = r.Run(":9090")
}
