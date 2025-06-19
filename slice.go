package main

import "fmt"

func main() {
	//arr := [6]string{"this", "is", "a", "go", "interview", "questions"}
	//fmt.Println(arr)
	//s := arr[1:4] // ["is", "a", "go"]
	// ptr = 1 reference of is
	// len = 3
	// cap = 5
	//fmt.Println(s)
	// slice to slice
	//s1 := s[1:2] // ["a"]
	//ptr = 1 reference of a
	//len = 1
	//cap = 4
	//fmt.Println(s1)
	//fmt.Println(len(s1))
	//fmt.Println(cap(s1))
	//s := []int{1, 2, 5} //slice literal
	//fmt.Println("slice", s, "len:", len(s), "cap:", cap(s))

	// slice delclear with make() function
	//s := make([]int, 3)
	//s[0] = 1
	//s[1] = 2
	//fmt.Println(s)
	//fmt.Println("len", len(s), "cap", cap(s))

	s := make([]int, 3, 5)
	s[2] = 10
	fmt.Println("len", len(s), "cap", cap(s))
	fmt.Println(s)

}
