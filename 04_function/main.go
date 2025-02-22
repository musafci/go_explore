package main

import "fmt"

// Example 1: Function with no return value

/*
func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

func main() {
	a := 10
	b := 20

	add(a, b)
	add(100, 150)
}
*/

// Example 2: Function with return value

/*
func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func main() {
	a := 10
	b := 20

	sum := add(a, b)
	fmt.Println(sum)
}
*/

// Example 3: Function with multiple return values

/*
func getValues(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	multi := num1 * num2

	return sum, multi
}


func main() {
	a := 10
	b := 20

	sum, multi := getValues(a, b)

	fmt.Println(sum)
	fmt.Println(multi)
}
*/

// Example 4: Function with named return values

/*
func getValues(num1 int, num2 int) (sum int, multi int) {
	sum = num1 + num2
	multi = num1 * num2

	return
}

func main() {
	a := 10
	b := 20

	sum, multi := getValues(a, b)

	fmt.Println(sum)
	fmt.Println(multi)
}
*/

// Example 5: Function with string parameter

/*
func getString(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	getString("John")
}
*/

// Example 6: Function with input from user

/*
func main() {
	fmt.Println("Welcome to my application!")

	fmt.Println("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	var num1 int
	var num2 int
	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)

	var sum int
	sum = num1 + num2

	fmt.Println("Your name is: ", name)
	fmt.Println("Summation is:", sum)

	fmt.Println("Goodbuy")
}
*/

// Example 7: Apply single responsibility principle in example 6

func welcome() {
	fmt.Println("Welcome to my application!")
}

func getUserName() string {
	fmt.Println("Enter your name: ")
	var name string
	fmt.Scanln(&name)
	return name
}

func getNumbers() (int, int) {
	var num1 int
	var num2 int

	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)

	return num1, num2
}

func getSummation(num1 int, num2 int) int {
	var sum int
	sum = num1 + num2
	return sum
}

func display(name string, sum int) {
	fmt.Println("Your name is: ", name)
	fmt.Println("Summation is:", sum)
}

func goodbuy() {
	fmt.Println("Goodbuy")
}

func main() {
	welcome()
	name := getUserName()
	num1, num2 := getNumbers()
	sum := getSummation(num1, num2)
	display(name, sum)
	goodbuy();
}