/*
	Passing Anonymous Functions as Arguments
	Since functions can be passed as arguments, anonymous functions are useful for higher-order functions.
*/

package main

import "fmt"

func operate(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func main() {
	result := operate(10, 20, func(x int, y int) int {  // this is also called function statement or function expression
		return x + y
	})
	fmt.Println(result)
}