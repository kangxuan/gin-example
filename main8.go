package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通过标记来反射
type message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Age     int    `json:"age"`
}

// 演示返回JSON数据两种方式
func main() {
	r := gin.Default()

	// 通过gin.H临时返回JSON数据
	r.GET("/main8/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "Shanla",
			"message": "hello golang!",
			"age":     18,
		})
	})

	// 通过结构体返回JSON数据
	r.GET("/main8/struct_json", func(c *gin.Context) {
		data := message{
			"小康",
			"你好，狗语言",
			18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
