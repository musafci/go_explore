package main

import "fmt"

var a int = 10

func main() {
	var age int = 30

	if age > 18 {
		var a int = 47 // variable shadowing
		fmt.Println(a)
	}

	fmt.Println(a)
}