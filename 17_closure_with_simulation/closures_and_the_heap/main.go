/*
	In Go, closures store captured variables in the heap memory instead of the stack, so they persist beyond their original function.
	Example with a Global Variable
*/

package main

import "fmt"

var x = 10

func createClosure() func() {
	y := 20 // Captured variable

	return func() {
		x++ // Modifies global x
		y++ // Modifies closure variable
		fmt.Println("x:", x, "y:", y)
	}
}

func main() {
	closure1 := createClosure()
	closure1() // Output: x: 11, y: 21
	closure1() // Output: x: 12, y: 22

	closure2 := createClosure()
	closure2() // Output: x: 13, y: 21 (new y for closure2)
}


/*
	Key Observations
	x is global and shared across all closures.
	y is specific to each closure and retains its own state.
*/