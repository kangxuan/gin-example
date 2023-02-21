package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 演示参数绑定
func main() {
	r := gin.Default()

	r.GET("/main12/user", func(c *gin.Context) {
		var u UserInfo
		// 进行参数绑定，如果是 GET 请求，只使用 Form 绑定引擎（query）
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Printf("userinfo:%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	})

	r.POST("/main12/user/form", func(c *gin.Context) {
		var u UserInfo
		// POST FORM-data请求依赖 Form 绑定引擎
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	})

	r.POST("/main12/user/json", func(c *gin.Context) {
		var u UserInfo
		// POST JSON请求依赖 JSON 绑定引擎，但经过测试，Tag中不添加json也能拿到数据
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	})

	_ = r.Run(":9090")
}
