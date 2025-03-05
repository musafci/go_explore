/*
	A closure in Go is an anonymous function that captures and remembers variables from its surrounding scope, 
	even after the scope has finished executing. 
	Closures allow functions to have state by keeping access to variables outside their own scope.
*/

package main

import "fmt"

const a = 10
var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age=", age)
	
	show := func() {
		money = money + a + p
		fmt.Println("Money=", money)
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("===Bank===")
}