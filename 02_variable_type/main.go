package main

import "fmt"

func main() {
	// var a int = 10
	a := 10 // type inference is used here to determine the type of a variable 
	a = 20 // allowed because a is already defined as int
	//a = "Hello" //not allowed because a is already defined as int

	var b float32 = 3.499
	b = 3.2222 // allowed because b is already defined as float32
	// b = "Hello" //not allowed because b is already defined as float32

	var c string = "Hello"
	// c = 123 //not allowed because c is already defined as string

	const pi = 3.1416
	// pi = 3.14159265358979323846264338327950288419716939937510 //not allowed because pi is already defined as constant

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(pi)
}
