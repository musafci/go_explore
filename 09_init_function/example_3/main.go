// Example 3: init in Multiple Files
// If a package has multiple files, all init functions run before main(), in the order they appear.

package main

import "fmt"

func init() {
	fmt.Println("I'm init function from main.go");
}

func main() {
	fmt.Println("I'm main function from main.go");
}