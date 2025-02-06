/*
	break
	default
	func
	interface
	select
	case
	defer
	go
	map
	struct
	chan 
	else
	goto
	package
	switch
	const
	fallthrough
	if
	range
	type
	continue
	for
	import
	return
	var
*/


package main

import "fmt"

func main() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
	// varibel
	var variableName type = string

	var a,b,c int

	var f,g,h type = value

	var n1,n2,n3 type = v1,v2,v3
	
	n4,n5,n6 := v1,v2,v3 // 简短声明，必须在函数内部。函数外面用war

	_, b := 34,35

	// const

	const constName = value

	const pi float32 = 3.145

	const Pi = 3.1415926
	const i = 10000
	const MaxThread = 10
	const prefix = "astaxie_"

	// built-in type
	var isnot bool
	var enabled,disabled = true,false
	func test(){
		var avai bool // 一般声明
		valid := false // 简短声明
		avai = true 
	}

	// number

	var c complex64 = 5 +5i
	fmt.Printf("Value is:%v",c)

	// string
	var hello string
	var empty string = ""
	func test(){
		a,b,c := "a","b","c"
		ja := "aas"
		fe = "aaa"
	} 

	// string can't change dircetly
	var s string = "hello"
	s[0] = 'c' // error

	s :=hello 
	c :=[]byte(s)
	c[0] = 'c'
	s2 := stirng(c)
	fmt.Printf("%s\n",s2)

	// string +

	s := "hello"
	s2 := "world"
	s3 := s + s2
	fmt.Printf("%s\n",s3)

	// change
	s := "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n",s)

	// multiple lines
	m := `hello
			world`
	
	// error
	error := errors.New("1111")
	if err != nil{
		fmt.Print(err)
	}

	// declare by group
	import(
		//
		//
		//
	)

	const(
		i =1 
		a=1
	)
	var(
		i int
		pi float32
	)
	// iota 枚举，没增加一行，const + 1.遇到新的const，归零

	// 大写字母可以导出

	//array
	var arr [n]type
	var arr [10]int
	arr[0] = 42

	a:= [3] int{1,2,3}
	b := [4]int(1,2,3)
	c := [...]int{4,5,6}

	// 嵌套数组
	doubleArray := [2][4]int{[4]{1,2,3,4},[4]{5,6,7,8}}

	easyArray := [2][4]int{{1,2,3,4},{5,6,7,8}}

	// slice 
	var fslice []int

	slice := []byte{'a','b','c','d','e'}

	var a,b []byte

	a = ar[2:5]
	//	ar[:n] === ar[0:n] 
	//  ar[n:]等价于ar[n:len(ar)]

	// map
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["ten"]= 10

	// map是引用类型
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"
	// m变化，m1也变化

	// slice = 一个指向数据，长度和容量的结构题的指针
	// make 和 new区别大同小异

	

}

