package main

import "fmt"

// for:- only construct for the loop in go

func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 5; i++ {
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

	// range
	for i := range 3 {
		fmt.Println(i)
	}

	// // infinite loop
	// for {
	// 	fmt.Println("hello")
	// }
}
