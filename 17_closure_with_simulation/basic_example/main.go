/*
	Basic example of closure
*/

package main

import "fmt"

func counter() func() int {
	count := 0 // Enclosed variable

	return func() int { // Anonymous function (closure)
		count++ // Increments count
		return count
	}
}

func main() {
	increment := counter() // Assign closure to variable

	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3

	newCounter := counter() // A new independent closure
	fmt.Println(newCounter()) // Output: 1
}


/*
	Explanation
	counter() returns an anonymous function that captures count.
	Each time increment() is called, it retains the last value of count and increases it.
	A new closure (newCounter()) starts with a fresh copy of count.
*/