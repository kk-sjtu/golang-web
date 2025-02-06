package main

import "fmt"

// type child struct{
// 	like string
// 	happy bool
// }

// type person struct{
// 	name string
// 	age int
// 	child
// }

type person struct{
	name string
	age int
}

type Human struct{
	name string
	age int
	phone string
}

type student struct{
	Human
	happy bool
	phone string
}




func main(){

	

	var p person
	p.name = "kk"
	p.age = 18
	fmt.Printf("the %s's age is %d\n",p.name,p.age)

	oneper := person{"k1",16}
	secper := person{"k2",22}

	oderman,res := older(oneper,secper)
	fmt.Printf("the older is %s,res is %d\n",oderman.name,res)

	Bob := student{
		Human:Human{"Bob",22,"111-222-333"},
		happy:true,
		phone:"222-333-444", // access first
	}
	fmt.Println("Bob's phone is:",Bob.phone)

}

func older(p1,p2 person)(person,int){
	if p1.age>p2.age{
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}