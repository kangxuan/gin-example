package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// template3 演示嵌套模板
func template3(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 嵌套模板解析需要将被包含文件放在后面
	t, err := template.ParseFiles("./template3.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Printf("解析模板文件失败，错误：%v", err)
		return
	}
	// 渲染模板
	name := "Shanla"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/template3", template3)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("The ListenAndServe Error: %v", err)
		return
	}

}
