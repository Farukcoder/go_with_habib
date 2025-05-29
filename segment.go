package main

import "fmt"

const e = 10

var p = 100

func calls() {
	Adds := func(x, y int) {
		z := x + y
		fmt.Println(z)
	}

	Adds(10, 50)
	Adds(e, 100)
}

//func Adder(x int, y int) {
//	z := x + y
//	fmt.Println(z)
//}

func main() {

	// code segment -> all function are go to // code segment
	// data segment -> all variable are go to data segment
	// stack -> all function call are go to stack segment allocate stake frame
	// heap
	//Adder(5, 2)
	//Adder(aa, 20)

	calls()
	fmt.Println(p)
}

func init() {
	fmt.Println("Hello")
}
