package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm() // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出入服务器的打印信息
	fmt.Println("path " ,r.URL.Path)
	fmt.Println("scheme " , r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key:" ,k)
		fmt.Println("value",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello vic")

}

func main()  {
	http.HandleFunc("/",sayHelloName)//设置访问路由
	err := http.ListenAndServe(":9090",nil)// 设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: " ,err)
	}

}
