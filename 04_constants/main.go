package main

import "fmt"

// short-hand does not work outside the func

func main() {
	const name string = "Hello"
	fmt.Println(name)

	const (
		port int = 5000
		host string = "localhost"
	)

	fmt.Println(port, host)
}