package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 演示接收Query string参数
func main() {
	r := gin.Default()

	r.GET("/main9", func(c *gin.Context) {
		// 获取url ? 之后的参数数据，叫query string
		// 第一种：通过query，key不存在会返回空字符串
		name := c.Query("name")
		// 第二种：通过DefaultQuery，key不存在会默认一个字符串
		gender := c.DefaultQuery("gender", "男")
		// 第三种：通过GetQuery,key不存在时会返回一个bool类型
		age, ok := c.GetQuery("age")
		if !ok {
			fmt.Println("age 不存在")
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"gender": gender,
			"age":    age,
		})
	})

	_ = r.Run(":9090")
}
