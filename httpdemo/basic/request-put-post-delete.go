package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
)
func get(){
	//r,err := http.Get("http://www.baidu.com")
	r,err := http.Get("https://httpbin.org/get")
	if err != nil{
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()

	content,err := ioutil.ReadAll(r.Body) // 读取响应体
	if err != nil{
		panic(err)
	}

	fmt.Println(string(content))
}

func post(){
	//r,err := http.Get("http://www.baidu.com")
	r,err := http.Post("https://httpbin.org/post","",nil)
	if err != nil{
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()

	content,err := ioutil.ReadAll(r.Body) // 读取响应体
	if err != nil{
		panic(err)
	}

	fmt.Println(string(content))
}

func put(){
	request ,err := http.NewRequest(http.MethodPut,"http://httpbin.org/put",nil)
	if err != nil{
		panic(err)
	}
	r,err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}

	defer func() {
		_ = r.Body.Close()
	}()

	content,err := ioutil.ReadAll(r.Body) // 读取响应体
	if err != nil{
		panic(err)
	}

	fmt.Println(string(content))

}

func delete(){
	request ,err := http.NewRequest(http.MethodDelete,"http://httpbin.org/delete",nil)
	if err != nil{
		panic(err)
	}
	r,err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}

	defer func() {
		_ = r.Body.Close()
	}()

	content,err := ioutil.ReadAll(r.Body) // 读取响应体
	if err != nil{
		panic(err)
	}

	fmt.Println(string(content))

}

func main(){
	//get()
	//post()
	//put()
	delete()
}