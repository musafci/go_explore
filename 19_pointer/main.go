package main

import "fmt"

type User struct {
	Name string
	Address string
}

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

	fmt.Println("--------------------------------")

	// Pointer with Function (Pass by Reference)
	fmt.Println("Pointer with Function (Pass by Reference)")
	num := 10
	increment(&num)
	fmt.Println("Num after increment:", num) // 11

	fmt.Println("--------------------------------")

	// Pointer to Struct
	fmt.Println("Pointer to Struct")
	user := User{Name: "Musa", Address: "123 Main St"}
	ptr := &user

	fmt.Println(ptr) // &{Musa 123 Main St}
	fmt.Println(*ptr) // {Musa 123 Main St}

	fmt.Println("Name:", ptr.Name) // Musa
	fmt.Println("Address:", ptr.Address) // 123 Main St

	ptr.Name = "John"
	ptr.Address = "456 Main St"
	fmt.Println(ptr) // &{John 456 Main St}
	fmt.Println(*ptr) // {John 456 Main St}

	fmt.Println("Name:", ptr.Name) // John
	fmt.Println("Address:", ptr.Address) // 456 Main St

	fmt.Println("--------------------------------")

	// Pointer to Array
	fmt.Println("Pointer to Array")
	numbers := [3]int{1, 2, 3}
	ptrArray := &numbers

	fmt.Println(ptrArray) // &[1 2 3]
	fmt.Println(*ptrArray) // [1 2 3]
	fmt.Println(ptrArray[0]) // 1
	fmt.Println(ptrArray[1]) // 2
	fmt.Println(ptrArray[2]) // 3

	ptrArray[0] = 4
	fmt.Println(numbers) // [4 2 3]


	fmt.Println("--------------------------------")

	// Nil Pointer
	fmt.Println("Nil Pointer")
	var ptrNil *int
    fmt.Println("Pointer p:", ptrNil) // <nil>

    if ptrNil == nil {
        fmt.Println("ptrNil is nil")
    }

	fmt.Println("--------------------------------")
	//Pointer to Pointer
	fmt.Println("Pointer to Pointer")
	num2 := 10
	ptr2 := &num2
	ptr3 := &ptr2
	ptr4 := &ptr3

	fmt.Println(num2) // 10
	fmt.Println(*ptr2) // 10
	fmt.Println(**ptr3) // 10
	fmt.Println(***ptr4) // 10
	

	fmt.Println("--------------------------------")
	//Pointer to Function
	fmt.Println("Pointer to Function")
	
	func(num *int) {
		*num++
	}(&num2)

	fmt.Println(num2) // 11

	fmt.Println("--------------------------------")

	// Pointer to Slice
	fmt.Println("Pointer to Slice")
	slice := []int{1, 2, 3}
	ptrSlice := &slice

	fmt.Println(slice) // [1 2 3]
	fmt.Println(ptrSlice) // &[1 2 3]
	fmt.Println(*ptrSlice) // [1 2 3]

	(*ptrSlice)[0] = 4
	fmt.Println(slice) // [4 2 3]

	fmt.Println("--------------------------------")
	
}


func increment(num *int) {
	*num++
}