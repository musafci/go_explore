package main

import "fmt"

var a int = 10
var b int = 20

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	p := 20
	q := 30

	add(p, q) // Will execute this line
	add(p, a) // Will execute this line
	add(q, b) // Will execute this line
	add(a, b) // Will execute this line
	//add(p, z) // Will not execute this line and get error because z is not defined in local scope or global scope
}
