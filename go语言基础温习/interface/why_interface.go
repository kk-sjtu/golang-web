package main

import "fmt"

// 接口是一种类型，是一种抽象的类型

type dog struct{}
type person struct{
	name string
}
func (d dog) say(){
	fmt.Println("dog say:汪汪汪")
}

type cat struct{}

func (c cat) say(){
	fmt.Println("cat say:喵喵喵")
}

func(p person)say(){
	fmt.Println("person say:啊啊啊")
}
// 定义一个类型，一个抽象的类型。只要实现了say这个方法的类型，
// 都可以成为sayer类型

type sayer interface{
	say()
}



// 接口不管你是什么类型，只管你要实现什么方法
func da(arg sayer){
	arg.say()
}

func main(){
	// c1 := cat{}
	// da(c1)
	// d1 := dog{}
	// da(d1)

	// p1 := person{name:"kk"}
	// da(p1)

	// 接口可以存储变量
	var s sayer
	c2 := cat{}
	s = c2
	fmt.Println(s)
	p2 :=person{name:"kk"}
	s = p2
	fmt.Println(s)
	d2 := dog{}
	s = d2
	fmt.Println(s)


}