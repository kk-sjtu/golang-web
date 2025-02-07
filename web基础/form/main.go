// main.go
package main

import (
    "fmt"
    "html/template"
    "net/http"
	"log"
)

func login(w http.ResponseWriter, r *http.Request) {
    log.Println("处理 /login 请求")
    if r.Method == "GET" {
        log.Println("处理 GET 请求")
        t, err := template.ParseFiles("login.gtpl")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        t.Execute(w, nil)
    } else {
        log.Println("处理 POST 请求")
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")

        log.Printf("用户名: %s\n", username)
        log.Printf("密码: %s\n", password)

        fmt.Fprintf(w, "用户名: %s\n", username)
        fmt.Fprintf(w, "密码: %s\n", password)
    }
}

func main() {
    http.HandleFunc("/login", login)
    fmt.Println("服务器启动，监听端口: 8080")
    http.ListenAndServe(":8080", nil)
}