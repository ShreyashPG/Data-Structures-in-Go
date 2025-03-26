package main

import (
	"fmt"
)

type person struct {
	name string
	age int
	next *person 
}

func (p * person) print(){
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.next)
}

func great(a int ) bool {
	if a > 10 {
		return true
	}
	return false
}

func main(){
		p1 := person{name:"Shreyash", age:25, next:nil}
	// 	fmt.Println(p1)
	// 	fmt.Println(p1.name)
	// 	fmt.Println(p1.age)
		// fmt.Println(p1.next)

		p2:= person{name:"Onkar", age:30, next:&p1}
		
		fmt.Println(p2)
	// 	fmt.Println(p2.name)
	// 	fmt.Println(p2.age)
			

	// 	fmt.Println("Printing next of p2")
    //    fmt.Println(p2.next.name)
	//    fmt.Println(p2.next.age)
	//    fmt.Println(p2.next.next)

	   	p2.print()

		flag :=great(10)
		fmt.Println(flag)
		

		var username string = "wagslane"
		var password int = 20947382822
	
		// don't edit below this line
		fmt.Println("Authorization: Basic", username, ":", password)
}
