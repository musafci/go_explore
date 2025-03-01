/*
	Go internal memory
	Go is designed with automatic memory management, including garbage collection, stack allocation, and heap allocation 
	to optimize performance prevent memory leak and memory overflow.

	There are four segments in Go memory:
	1. Code segment: stores compiled instructions, functions definitions, and constants
	2. Data segment: stores global variables and static variables
	3. Heap segment: stores functions calls & local variables 
	4. Stack segment: stores dynamically allocated memory, slice, maps, structs, and channels.
*/


package main

import "fmt"

var a = 10

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	add(4, 7)
	add(a, 7)
}

func init() {
	fmt.Println("This is init function")
}