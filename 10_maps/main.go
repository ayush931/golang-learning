package main

import (
	"fmt"
	"maps"
)

// map -> hash, object, dict
func main() {
	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// Imp:- If the key does not exists, it will return the 0 value

	// getting the element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["hello"]) // empty string

	m1 := make(map[string]int)
	m1["age"] = 24
	m1["score"] = 10
	fmt.Println(m1["age"])
	fmt.Println(m1["hello"]) // will return 0 as no key exists

	fmt.Println(len(m1))

	fmt.Println(m1)
	delete(m1, "score")
	fmt.Println(m1)

	clear(m1)
	fmt.Println(m1)

	m2 := map[string]int{"price": 20, "phone": 3}
	fmt.Println(m2)

	// _, ok := m["price"]
	_, ok := m2["price"]

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	v, ok := m2["hello"] // if element not exists then 0
	fmt.Println(v)

	m3 := map[string]int{"price": 20, "phone": 3}
	m4 := map[string]int{"price": 20, "phone": 3}

	fmt.Println(maps.Equal(m3, m4)) // boolean result
}
