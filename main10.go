package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示获取Form参数
func main() {
	c := gin.Default()

	// 解析模板
	//c.LoadHTMLFiles("./main10Templates/login.tmpl", "./main10Templates/index.tmpl")
	c.LoadHTMLGlob("./main10Templates/*")
	c.GET("/main10/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil) // 渲染模板
	})

	c.POST("/main10/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"username": username,
			"password": password,
		})
	})

	_ = c.Run(":9090")
}
