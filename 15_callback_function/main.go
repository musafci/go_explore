/*
	Callback Function:
	- A function that takes one or more functions as arguments or returns a function as its result.
*/

package main

import "fmt"

func call() func(x int, y int) int { // this is callback function
	return add
}

func add(x int, y int) int {
	return x + y
}

func main() {
	// fmt.Println(call()(10, 20))
	sum := call()(10, 20)
	fmt.Println(sum)
}
