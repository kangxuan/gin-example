package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 通过Any处理不同的请求方式，相当于一个大杂烩
	r.Any("/main15/any", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"method": "Any",
			})
		}
	})

	// 路由组，多用于接口多版本
	v1 := r.Group("/main15/v1")
	{
		v1.POST("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "add user",
			})
		})

		v1.PUT("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "update user",
			})
		})

		v1.DELETE("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete user",
			})
		})

		v1.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get a user",
			})
		})

		// 嵌套路由组
		xx := v1.Group("/xx")
		{
			xx.GET("/info", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "xx info",
				})
			})
		}
	}

	// 未定义路由的走这里的逻辑
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found!",
		})
	})
	_ = r.Run(":9090")
}
