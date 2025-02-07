package main

import "time"


//func main() {
// 	c := make(chan int)
// 	o := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 				case c <- 1:
					
// 				case <- time.After(5 * time.Second):
// 					println("timeout")
// 					o <- true
// 					break
// 			}
// 		}
// 	}()

// 	go func(){
// 		for{
// 			fmt.Println(<-c)
// 		}
// 	}()

// 	<- o
// }


func main() {
	c := make(chan int)
	o := make(chan bool)
	c <- 1
	go func() {
		for {
			select {
				case v := <- c:
					println(v)
				case <- time.After(5 * time.Second):
					println("timeout")
					o <- true
					break
			}
		}
	}()
	<- o
}

// 如果你在主函数中发送数据到通道 c，
// 并且在一个 goroutine 中接收数据，
// 仍然会出现死锁的原因可能是因为接收操作没有及时开始，
// 或者在发送操作之前没有启动接收操作。