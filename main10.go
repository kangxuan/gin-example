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
		// 第一种：通过PostForm 不存在则是空字符串
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 第二种，通过DefaultPostForm 不存在返回默认值
		//username := c.DefaultPostForm("username", 'xxx')

		// 第三种，通过GetPostForm 不存在返回false
		//username, ok := c.GetPostForm("username")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"username": username,
			"password": password,
		})
	})

	_ = c.Run(":9090")
}
