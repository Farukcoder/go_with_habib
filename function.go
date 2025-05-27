package main

import "fmt"

var a = 10

func Add(x int, y int) {
	fmt.Println(x + y)
}

func main() {
	//standard function
	//Add(10, 20)

	//init function
	//1.no call this function automatically call this
	fmt.Println("Hello Init Function")
	fmt.Println(a)
}

func init() {
	fmt.Println(a)
	a = 20
}
