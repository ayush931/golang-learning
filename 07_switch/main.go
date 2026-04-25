package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	var i int = 2

	switch i {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")

	case 4:
		fmt.Println("Four")

	default:
		fmt.Println("Other")
	}

	// multiple switch case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend day")

	default:
		fmt.Println("Working day")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an Integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("Other", t)
		}
	}

	whoAmI(true)
}