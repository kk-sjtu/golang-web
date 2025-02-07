package main

import(
	"net/http"
	"io/ioutil"
	"fmt"
	"net/url"
)
func printBody(r *http.Response){
	content,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s",content)

}

func requestByParams(){
	request,err := http.NewRequest(http.MethodGet,"https://httpbin.org/get",nil)
	if err != nil{
		panic(err)
	}
	params := make(url.Values)
	params.Add("name","kk")
	params.Add("age","18")
	fmt.Println(params.Encode())
	request.URL.RawQuery = params.Encode()

	r,err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()
	printBody(r)

}

func requestByHeader(){
	request,err := http.NewRequest(http.MethodGet,"https://httpbin.org/get",nil)
	if err != nil{
		panic(err)
	}
	request.Header.Add("User-Agent","kk")
	r,err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()
	printBody(r)
	}
func main(){
	// 1. 如何设置请求的查询参数
	requestByParams()

	// 2. 如何定制请求头，比如修改user-agent
	requestByHeader()

	// user agent反扒策略？
}