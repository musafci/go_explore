//Example 2: Using init for Global Variable Initialization

package main

import "fmt"

var config string

func init() {
    config = "Application Configured"
    fmt.Println("Configuring application...")
}

func main() {
    fmt.Println(config) // Output: Application Configured
}