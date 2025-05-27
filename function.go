package main

import "fmt"

//var a = 10

//func Add(x int, y int) {
//	fmt.Println(x + y)
//}

var adder = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	//standard function
	//Add(10, 20)

	//init function
	//1.no call this function automatically call this
	//fmt.Println("Hello Init Function")
	//fmt.Println(a)

	//iife Function anonymous function
	//immediately invoked Function Expression
	//func(a, b int) {
	//	c := a + b
	//	fmt.Println(c)
	//}(5, 7)

	//function expression not call up of course call down
	//add(2, 5)

	add := func(a, b int) {
		fmt.Println(a + b)
	}

	add(1, 2)
	adder(4, 5)
}

//func init() {
//	//fmt.Println(a)
//	//a = 20
//	fmt.Println("init")
//}
