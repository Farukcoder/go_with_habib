package main

import "fmt"

const ab = 10

var pp = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age = ", age)

	show := func() {
		money = money + ab + pp
		fmt.Println(money)
	}

	return show
}

func Call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	Call()
}

func init() {
	fmt.Println("===blank===")
}
