package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示重定向
func main() {
	r := gin.Default()

	r.GET("/main14/index", func(c *gin.Context) {
		// 跳转到外部链接
		c.Redirect(http.StatusMovedPermanently, "http://baidu.com")
	})

	r.GET("/main14/a", func(c *gin.Context) {
		c.Request.URL.Path = "/main14/b" // 将请求到URI修改
		r.HandleContext(c)               // 继续后续的处理
		// 下面的可以继续执行哦
		c.JSON(http.StatusOK, gin.H{
			"message": "ok1",
		})
	})

	r.GET("/main14/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	_ = r.Run(":9090")
}
