package main

import "fmt"

// golang does not have ternary operator, if else will be used.

func main() {
	const age int = 18

	if age >= 18 {
		fmt.Println("can vote")
	} else if age <= 5 {
		fmt.Println("Person is child")
	} else {
		fmt.Println("Cannot vote")
	}

	const user string = "user"
	const permission bool = true

	if user == "admin" && permission == true {
		fmt.Println("You are admin")
	} else if user == "user" && permission == true {
		fmt.Println("User account")
	} else {
		fmt.Println("Not permitted")
	}

	if signal := "Green"; signal == "red" {
		fmt.Println("STOP")
	} else if signal == "Yellow" {
		fmt.Println("WAIT")
	} else if signal == "Green" {
		fmt.Println("GO")
	} else {
		fmt.Println("Wrong")
	}
}