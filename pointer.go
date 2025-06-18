package main

import "fmt"

//func print(numbers *[3]int) {
//	fmt.Println(numbers)
//}

type User struct {
	Name   string
	Age    int
	Salary float64
	foods  []string
}

func main() {
	//var x = 20
	//fmt.Println(x)
	//addr := &x
	//*addr = 30
	//fmt.Println(x, addr)
	//fmt.Println(*addr)
	//fmt.Println(x)
	//arr := [3]int{1, 2, 3}
	//print(&arr)
	userInfo := User{
		Name:   "Faruk",
		Age:    20,
		Salary: 100000,
	}
	pointer := &userInfo
	fmt.Println(pointer.Name)
}
