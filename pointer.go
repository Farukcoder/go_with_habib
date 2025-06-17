package main

import "fmt"

func main() {
	var x = 20
	addr := &x
	fmt.Println(x, addr)
}
