package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./templates4/base.tmpl", "./templates4/index1.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles Error:%v", err)
		return
	}
	// 渲染模板
	name := "Shanla"
	// 因为定义了多个模板，具体渲染哪个模板需要指定
	err = t.ExecuteTemplate(w, "index1.tmpl", name)
	if err != nil {
		fmt.Printf("ExecuteTemplates Error:%v", err)
		return
	}
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./templates4/base.tmpl", "./templates4/index2.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles Error:%v", err)
		return
	}
	// 渲染模板
	name := "小康"
	// 因为定义了多个模板，具体渲染哪个模板需要指定
	err = t.ExecuteTemplate(w, "index2.tmpl", name)
	if err != nil {
		fmt.Printf("ExecuteTemplates Error:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/index1", index1)
	http.HandleFunc("/index2", index2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("The Listen And Serve Error:%v", err)
		return
	}
}
