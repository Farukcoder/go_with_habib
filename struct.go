package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user1 := User{
		Name: "Faiza",
		Age:  6,
	}

	fmt.Println("name", user1.Name)
	fmt.Println("age", user1.Age)

	user2 := User{
		Name: "Zarif",
		Age:  2,
	}

	fmt.Println("name", user2.Name)
	fmt.Println("age", user2.Age)
}
