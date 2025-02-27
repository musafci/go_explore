package main

import "fmt"

func add(x int, y int) int { // x and y are parameters
	return x + y
}

func main() {
	sum := add(10, 20) // 10 and 20 are arguments
	fmt.Println(sum)
}