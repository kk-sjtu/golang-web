package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func responseBody(r *http.Response){
	content,_ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s",content)
}

func header(r *http.Response){
	fmt.Println("header:",r.Header)
	fmt.Println("content-length:",r.Header.Get("Content-Length"))
	fmt.Println("content-type:",r.Header.Get("Content-Type"))
	
}

func status(r *http.Response){
	fmt.Println("status code:",r.StatusCode) // 状态码
	fmt.Println("status:",r.Status) // 状态描述

}

func encoding(){

}

func main(){
	r, err := http.Get("https://httpbin.org/get")
	if err != nil{
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()

	//responseBody(r) 
	//status(r)
	header(r)



}