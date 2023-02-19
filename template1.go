package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name, Gender string
	Age          int
}

func Index(w http.ResponseWriter, r *http.Request) {
	// 添加模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("ParseFile error:%v", err)
		return
	}
	// 解析模板
	u1 := User{
		Name:   "Shanla",
		Gender: "男",
		Age:    30,
	}
	m1 := map[string]interface{}{
		"name":   "小康",
		"gender": "男",
		"age":    18,
	}
	hobbyList := []string{
		"抽烟",
		"喝酒",
		"烫头",
	}
	err = t.Execute(w, map[string]interface{}{
		"u1":        u1,
		"m1":        m1,
		"hobbyList": hobbyList,
	})
	if err != nil {
		fmt.Printf("Execute Error:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", Index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("ListenAndServe err:%v", err)
		return
	}
}
