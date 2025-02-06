package main

import "fmt"

func isodd(num int) bool{
	if num%2 ==0{
		return false
	}
	return true
}

func iseven(num int)bool{
	if num%2 ==0{
		return true
	}
	return false
}

func filter(slice []int,f func(int)bool)[]int{
	news := []int{}
	for _,value := range slice{
		if f(value){
			news = append(news,value)
		}

	}
	return news
}

func main(){



	slice := []int{1,2,3,4,5,6,7}
	fmt.Println(filter(slice,isodd))
	fmt.Println(filter(slice,iseven))

}

