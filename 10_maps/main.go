package main

import "fmt"

//  map -> hash, object, dict
func main() {
	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// Imp:- If the key does not exists, it will return the 0 value

	// getting the element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["hello"])		// empty string

	m1 := make(map[string]int)
	m1["age"] = 24
	fmt.Println(m1["age"])

	fmt.Println(len(m))
}