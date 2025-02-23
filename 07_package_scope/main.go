package main

import "fmt"
import "myapp.dev/calculator"

func main() {
	fmt.Println("Hello, World!")
	
	/*
		below are the functions from the calculator.go file, to be used in the main package,
		you need to run "go run main.go" to run this file, no need to run "go run main.go calculator.go"
		because the calculator.go file is already imported in the main.go file
		and the functions are already defined in the calculator.go file
	*/
	calculator.Add(10, 20)
	calculator.Subtract(10, 20)
	calculator.Multiply(10, 20)
	calculator.Divide(10, 20)

	/*
		below is a function from the hello.go file, to be used in the main package,
		to run this file, you need to run "go run main.go hello.go", otherwise it will not work and you will get an error	
	*/
	printHello() 
}
