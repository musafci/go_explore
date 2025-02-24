package main

import "fmt"

func init() {
	fmt.Println("I'll be work first before main function.");
}
func main() {
	fmt.Println("Hello, World!")
}