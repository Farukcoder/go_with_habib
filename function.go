package main

import "fmt"

//var a = 10

//func Add(x int, y int) {
//	fmt.Println(x + y)
//}

//var adder = func(x, y int) {
//	fmt.Println(x + y)
//}

// higher order function
func processOrder(x int, y int, op func(p int, q int)) {
	op(x, y)
}

func Add(a int, b int) { //parameter
	c := a + b
	fmt.Println(c)
}

func call() func(int, int) {
	return Add
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

	//add := func(a, b int) {
	//	fmt.Println(a + b)
	//}
	//
	//add(1, 2)
	//adder(4, 5)

	//parameter vs argument
	//add(2, 5) //argument

	//first order function

	//higher order function
	processOrder(10, 20, Add) //passing function as argument
	sum := call()             // function expression
	sum(20, 20)

	//first class citizens (variable assign data)
	//call back function

}

//func init() {
//	//fmt.Println(a)
//	//a = 20
//	fmt.Println("init")
//}
