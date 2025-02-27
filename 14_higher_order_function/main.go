/*
	Higher Order Function:
	- A function that takes one or more functions as arguments or returns a function as its result.
*/

package main

import "fmt"

func operate(x int, y int, f func(int, int) int) int { // f is a function that takes two int arguments and returns an int
	return f(x, y)
}

func add(x int, y int) int {
	return x + y
}

func main() {
	result := operate(4, 7, add)
	fmt.Println(result)
}