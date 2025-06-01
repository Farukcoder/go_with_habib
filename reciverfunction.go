package main

import "fmt"

type Users struct {
	Name string
	Age  int
}

func printUserDetails(usr Users) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

// receiver function
func (usr Users) printUserDetails() {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

func (usr Users) call(a int) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println(a)
}

func main() {
	user1 := Users{
		Name: "Faiza",
		Age:  6,
	}

	//printUserDetails(user1)
	user1.printUserDetails()
	user1.call(2)
	user2 := Users{
		Name: "Zarif",
		Age:  2,
	}
	user2.printUserDetails()
	user2.call(3)
}
