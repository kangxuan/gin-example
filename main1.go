package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示Restful风格
func main() {
	r := gin.Default()
	// restful风格
	// 查
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "get a book",
		})
	})
	// 增
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "create a book",
		})
	})
	// 改
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "update a book",
		})
	})
	// 删
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete a book",
		})
	})

	r.Run()
}
