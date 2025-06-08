
🔹 Understanding Pointers in Go (Golang)
Just shared a basic pointer example for Go beginners 👇

This code demonstrates how:

A pointer stores the memory address of a variable.

You can access and modify the original value using the pointer.

Both the variable and pointer reference the same memory location.

📌 Key takeaway: Changing the value via the pointer also updates the original variable!

Happy Go coding! 💻
#Golang #Pointers #BackendDevelopment #GoLangTips #LearningByDoing

package main

import "fmt"

func main() {
	// Basic Pointer Example
	fmt.Println("Basic Pointer Example")
	var a int = 10 // a is a variable of type int
	var p *int = &a // p is a pointer to an int, and it is pointing to the address of a

	fmt.Println(a) // 10 (value of a)
	fmt.Println(p) // 0xc0000160a0 (address of a)
	fmt.Println(*p) // 10 (value of a)

	*p = 20 // change the value of a to 20
	fmt.Println(a) // 20 (value of a)
	fmt.Println(*p) // 20 (value of a)

	fmt.Println(p) // 0xc0000160a0 (address of a)
	fmt.Println(&a) // 0xc0000160a0 (address of a)

	fmt.Println(a) // 20 (value of a)
	fmt.Println(*p) // 20 (value of a)
}
