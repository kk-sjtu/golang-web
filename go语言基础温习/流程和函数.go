package main

import "fmt"

// if 可以声明变量
func main(){
	
	if x:= computedValue();x>10{
		fmt.Println("x is greater than 10")
	}else{
		fmt.Println("x is less than 10")
	}

	if interger == 3{
		fmt.Println("")
	}else if integer <3{
		
	}else{

	}

	sum:=0
	for index = 0;index <10;index++{
		sum += index
	}

	sum := 1
	for sum <1000{
		sum += sum
	}

	// switch最后相当于带有break

	func Sum(a int,b int)(s int){
		return a+b // 当然也可以返回多个值
	}
	a :=3
	b :=4
	sum = Sum(a,b)
	fmt.Printf("sum(%d,%d) = %d\n",a,b,sum)

	func add(a *int)int{
		*a = *a + 1
		return *a

	}
	x :=3
	x2 := add(&x)

	// 指针使多个函数能操作同一个对象
	// 指针比较轻量级


	func ReadWrite() bool {
		file.Open("file")
		defer file.Close()
		if failureX {
			return false
		}
		if failureY {
			return false
		}
		return true
	}
	// 或许可以理解为，里面不管干了什么，defer留着最后关门

	for i:=0;i<5;i++{
		defer fmt.Printf("%d ",i)
	}

	

}
