package main

import (
	"fmt"
)

type User struct {
	Name string
	Age int
}

func main() {
	// Example 1
	// user := User {
	// 	Name: "Md Abu Musa",
	// 	Age: 34,
	// }
	// fmt.Println("Name:", user.Name)
	// fmt.Println("Age:", user.Age)


	// Example 2
	// var user User
	// user.Name = "Md Abu Musa"
	// user.Age = 34

	// var user2 User
	// user2.Name = "Muhammad"
	// user2.Age = 2

	// fmt.Println("Name:", user.Name, "Age:", user.Age)
	// fmt.Println("Name:", user2.Name, "Age:", user2.Age)


	// Example 3
	var user User // just declaration of other variable like int, string, etc. example: var a int

	user = User {
		Name: "Musa",
		Age: 34,
	}

	fmt.Println(user.Name, user.Age)
}