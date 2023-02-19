package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// template2 演示如何自定义函数
func template2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t := template.New("template2.tmpl")
	// 定义一个自定义函数
	praise := func(name string) string {
		return name + "又年轻了"
	}
	// 传递一个函数给模板
	t.Funcs(template.FuncMap{
		"praise": praise,
	})
	_, err := t.ParseFiles("./template2.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles Error:%v", err)
	}
	// 执行模板
	name := "Shanla"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Execute Error:%v", err)
	}
}

func main() {
	http.HandleFunc("/template2", template2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("The ListenAndServe Error: %v", err)
		return
	}
}
