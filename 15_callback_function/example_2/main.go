package main

import "fmt"

// Function that accepts a callback
func execute(x int, y int, callback func(int, int) int) int {
	return callback(x, y)
}

// Add function
func add(x int, y int) int {
	return x + y
}

func main() {
	// Pass 'add' function as a callback
	result := execute(10, 20, add)
	fmt.Println(result) // Output: 30
}
