package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse files error:%v", err)
		return
	}
	// 渲染模板
	name := "Shanla"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Execute error:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("ListenAndServer error:%v", err)
		return
	}
}
