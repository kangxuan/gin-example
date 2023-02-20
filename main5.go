package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 演示自定义模板函数
func main() {
	r := gin.Default()
	// 定义模板函数，可以在模板中使用
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 定义模板
	// 解析模板
	r.LoadHTMLFiles("main5Templates/index.tmpl")
	r.GET("/main5/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://baidu.com'>百度</a>") // 渲染模板
	})
	_ = r.Run(":9090")
}
