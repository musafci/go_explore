package main

import "fmt"

var a int = 10
var b int = 20

func main() {
	age := 18

	if age >= 18 {
		num := 10
		fmt.Println("I am", age, "years boy!")
		fmt.Println("I have a younger brother, his age is", num)
	}

	// fmt.Println(num) // Get error because num is not defined in the local scope or global scope
}
