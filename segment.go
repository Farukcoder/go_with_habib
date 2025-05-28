package main

import "fmt"

var a = 10

func Adder(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {

	// code segment -> all function are go to // code segment
	// data segment -> all variable are go to data segment
	// stack -> all function call are go to stack segment allocate stake frame
	// heap
	Adder(5, 2)
	Adder(a, 20)

}

func init() {
	fmt.Println("Hello")
}
