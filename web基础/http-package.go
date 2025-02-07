
package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

// 一个Web服务器的主要工作流程可以简化为以下几个步骤：

// 监听端口：服务器需要监听一个特定的端口，等待客户端的请求。
// 接收请求：当客户端发起请求时，服务器接收并处理该请求。
// 处理请求：服务器根据请求的内容，调用相应的处理函数（handler）来生成响应。
// 返回响应：服务器将处理结果返回给客户端。

// Request：用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、url等信息
// Conn: 用户的每次请求链接信息
// Response: 服务器需要反馈给客户端的信息
// Handler: 处理请求和生成返回信息的处理逻辑


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
