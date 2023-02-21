package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示接收path参数
func main() {
	r := gin.Default()

	r.GET("/main11/user/:username/:age", func(c *gin.Context) {
		// 通过Param获取path参数
		username := c.Param("username")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
		})
	})

	_ = r.Run(":9090")
}
