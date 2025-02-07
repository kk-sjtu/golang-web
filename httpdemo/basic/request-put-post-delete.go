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



// POST 方法
// 语义：

// POST 方法用于向服务器提交数据，通常用于创建资源。
// POST 请求会导致服务器在目标资源下创建一个新的子资源。
// 使用场景：

// 提交表单数据。
// 上传文件。
// 发送数据以进行处理（例如，提交评论、发布文章）。
// 特点：

// POST 请求是非幂等的（non-idempotent），即多次相同的 POST 请求可能会导致不同的结果。例如，提交两次相同的表单可能会创建两个不同的资源。
// POST 请求的响应通常包含新创建资源的描述或状态。
// PUT 方法
// 语义：

// PUT 方法用于向服务器上传数据，通常用于更新资源。
// PUT 请求会在目标资源位置上创建或替换资源。
// 使用场景：

// 更新现有资源的全部内容。
// 如果资源不存在，可以创建资源（取决于服务器实现）。

// 特点：

// PUT 请求是幂等的（idempotent），即多次相同的 PUT 请求会产生相同的结果。例如，多次上传相同的数据会覆盖之前的数据，不会创建多个资源。
// PUT 请求的响应通常包含更新后资源的描述或状态。
// 示例代码中的 POST 和 PUT 请求
