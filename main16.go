package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// m1 定义一个中间件m1
func m1(c *gin.Context) {
	log.Println("m1 in...")
	c.Next()
	//c.Abort() // 阻止调用后续的处理函数
	log.Println("m1 out...")
}

// StatCost 定义一个接口耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Cost in ..")
		start := time.Now()
		c.Set("name", "小王子")
		c.Next() // 先调用后续的处理函数
		//c.Abort() // 阻止调用后续的处理函数，包括后面的中间件
		cost := time.Since(start)
		log.Println(cost)
		//log.Println(c.Get("name"))
		log.Println("Cost out ..")
	}
}

// AuthMiddleWare 权限验证中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断是否登录
		// 如果登录
		// c.Next
		// 否则
		// c.Abort
	}
}

func indexHandler(c *gin.Context) {
	log.Println("IndexHandler in..")
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	log.Println("IndexHandler out..")
}

func main() {
	r := gin.Default() // 默认使用了Logger和Recovery中间件
	//r:=gin.New() // 新建一个没有任何默认中间件的路由
	// 使用中间件
	r.Use(StatCost(), m1) // 全局注册中间件，可以注册多个中间件，按照顺序执行

	r.GET("/main16/index", indexHandler)

	_ = r.Run(":9090")
}
