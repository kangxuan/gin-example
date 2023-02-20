package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 定义模板
	// 解析模板
	r.LoadHTMLGlob("main4Templates/**/*") // 传入一个正则， ** - 目录 * - 文件

	r.GET("/main4/pos/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pos/index.tmpl", gin.H{ // 渲染模板
			"title": "这是pos/index的标题",
		})
	})

	r.GET("main4/user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title": "这是user/index的标题",
		})
	})

	_ = r.Run(":9090")
}
