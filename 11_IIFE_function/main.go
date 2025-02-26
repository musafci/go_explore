package main

import "fmt"

func main() {
	resutl := func(x int, y int) int { // this is also called function statement or function expression
		return x + y
	} (10, 20)

	fmt.Println(resutl)
}