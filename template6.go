package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// template6 演示被自动转义
func template6(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./template6.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles Error:%v", err)
		return
	}
	// 渲染模板
	t1 := "<script>alert(1234)</script>"
	t2 := "<a href='baidu.com'>百度</a>"
	err = t.Execute(w, map[string]interface{}{
		"t1": t1,
		"t2": t2,
	})
	if err != nil {
		fmt.Printf("Execute Error:%v", err)
		return
	}
}

func template6a(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 在解析模板之前先自定义一个函数到模板中
	t := template.New("template6a.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 解析模板
	t, err := t.ParseFiles("./template6a.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles Error:%v", err)
		return
	}
	// 渲染模板
	t1 := "<script>alert(1234)</script>"
	t2 := "<a href='baidu.com'>百度</a>"
	err = t.Execute(w, map[string]interface{}{
		"t1": t1,
		"t2": t2,
	})
	if err != nil {
		fmt.Printf("Execute Error:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/template6", template6)
	http.HandleFunc("/template6a", template6a)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("ListenAndServe Error:%v", err)
		return
	}
}
