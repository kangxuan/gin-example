package routers

import (
	controllers2 "gin-example/bubble/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter 注册路由
func SetupRouter() *gin.Engine {
	// 实例化gin实例
	r := gin.Default()

	// 静态文件
	r.Static("/static", "./static")
	// 解析模板
	r.LoadHTMLFiles("./templates/index.html")

	r.GET("/index", controllers2.IndexHandler)

	v1 := r.Group("/v1")
	{

		v1.GET("/todo", controllers2.GetGotoList)
		v1.POST("/todo", controllers2.AddGoto)
		v1.PUT("/todo/:id", controllers2.UpdateGoto)
		v1.DELETE("/todo/:id", controllers2.DeleteGoto)
	}

	return r
}
