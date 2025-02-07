package main

import "fmt"

func sum(a []int, c chan int){
	total :=0
	for _,v := range a{
		total += v
	}
	c <- total
}

func main(){
	a := []int{7,2,8,-9,4,0}
	
	c := make(chan int)

	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)

	x,y := <-c,<-c 
	// 后进先出
	// 两个goroutine，一个计算前一半，一个计算后一半
	// 两个goroutine计算完后，main goroutine计算两个goroutine的结果
	

	fmt.Println(x,y,x+y)

}