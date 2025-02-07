
package main

import "fmt"

// func main() {
// 	c := make(chan int, 1)//修改2为1就报错，修改2为3可以正常运行
// 	c <- 1
// 	c <- 2
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// }
//         //修改为1报如下的错误:
//         //fatal error: all goroutines are asleep - deadlock!
// 		// 实际上就是死锁

func fibonacci(n int,c chan int){
	x,y := 0,1
	for i:=0;i<n;i++{
		c <- x
		x,y = y,x+y
	}
	close(c)
}

func main(){
	c := make(chan int,10)
	go fibonacci(cap(c),c)
	for i:= range c{
		fmt.Println(i)
	}
	v,ok := <-c
	fmt.Println(v,ok)
}