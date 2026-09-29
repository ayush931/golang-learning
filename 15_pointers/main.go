package main

import "fmt"

// by value
func changeName(num int) {
	num = 5
	fmt.Println("in changeNum", num)
}

// by reference
func changeNums(num *int) {
	*num = 5
	fmt.Println("In changeNums", *num)
}

func main() {
	num := 1
	changeName(num)

	fmt.Println("After changeNum in main", num)

	fmt.Println("Memory address", &num)

	changeNums(&num)

	fmt.Println("After changeNum in main", num)
}
