package main

import "fmt"

func main() {
	var name string = "first_name"
	fmt.Println(name)

	// auto infer the type
	var another_name = "second_name"
	fmt.Println(another_name)

	// short-hand
	names := "short_name"
	fmt.Println(names)

	var price float32 = 50.3
	fmt.Println(price)
}