package	main

import "fmt"

func webwork(){
	fmt.Println(
		"url",
		"浏览器请求DNS",
		"dns 获取域名对应ip",
		"ip找到ip对应的服务器",
		"TCp",
		"浏览器发送完http request请求",
		"服务器接收请求包之后处理请求包",
		"服务器调用自身服务",
		"服务器返回Http响应",
		"客户端收到服务器响应后开始渲染响应中的主体",
		"收到全部内容后断开连接",
	)
}
func main(){



	webwork()

}