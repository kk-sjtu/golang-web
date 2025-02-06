package main

import (
	"fmt"
	"math"
)

type figure struct{
	id int
	Circle
}

type Circle struct{
	r float32

}

type Rectangle struct{
	w float32
	l float32
}

func(c Circle)Area()float32{
	return c.r*c.r*math.Pi
}

func(rc Rectangle)Area()float32{
	return rc.w*rc.l
}

func main(){
	c := Circle{5}
	fmt.Println("Area of Circle(c)=",c.Area())

	rc := Rectangle{3,4}
	fmt.Println("Area of Rectangle(rc)=",rc.Area())

	figure := figure{1,Circle{5}}
	fmt.Println("Area of Circle(figure)=",figure.Area())
}

//下面代码用指针是必要的。(b *Box)，只有用指针才可以对原来对象进行修改。

// func (b *Box) SetColor(c Color) {
// 	b.color = c // 自动识别指针所指向的对象
// }



// func (bl BoxList) PaintItBlack() {
// 	for i := range bl {
// 		bl[i].SetColor(BLACK)
// 	}
// }