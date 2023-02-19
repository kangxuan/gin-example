package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// template5 演示如何修改默认的标识符
func template5(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("template5.tmpl").Delims("{[", "]}").ParseFiles("template5.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles Error:%v", err)
		return
	}
	// 渲染模板
	_ = t.Execute(w, "Shanla")
}

func main() {
	http.HandleFunc("/template5", template5)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("ListenAndServe Error:%v", err)
		return
	}
}
