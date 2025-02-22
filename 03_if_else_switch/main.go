package main

import "fmt"

func main() {
	/*
	age := 18

	if age > 18 {
		fmt.Println("You are an adult")
	} else if age == 18 {
		fmt.Println("You are 18 years old")
	} else {
		fmt.Println("You are a minor")
	}
	*/

	age := 18
	gender := "male"

	if age > 18 && gender == "male" {
		fmt.Println("You are an adult male")
	} else if age > 18 && gender == "female" {
		fmt.Println("You are an adult female")
	} else {
		fmt.Println("You are a minor")
	}


	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	case "Thursday":
		fmt.Println("Today is Thursday")
	}
}
