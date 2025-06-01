package main

import "fmt"

func main() {
	var arr [2]int

	arrs := [2]int{1, 2}

	arr[1] = 6
	arr[0] = 9

	fmt.Println(arr)
	fmt.Println(arrs)
}
