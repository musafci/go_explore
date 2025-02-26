package main

import "fmt"

func main() {
	add := func(x int, y int) int { // this is also called function statement or function expression
		return x + y
	}

	fmt.Println(add(10, 20))
}