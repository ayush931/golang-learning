package main

import (
	"fmt"
	"slices"
)

// dynamic
// most used constructor to go
// various method

func main() {
	// uninitialize slice is nil
	var nums [] int

	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var num = make([] int, 2, 5)
	fmt.Println(num == nil)  // false

	// capacity -> maximum number of element can fit
	fmt.Println(cap(num))
	
	fmt.Println(len(num))

	num = append(num, 5, 5, 4, 3, 2, 7, 8, 9, 1)
	fmt.Println(num)
	fmt.Println(cap(num))		// resize automatically as the ele

	var numbs = make([] int, 0, 5);
	numbs = append(numbs, 2)

	var numbs2 = make([] int, len(numbs))

	copy(numbs2, numbs);

	fmt.Println(numbs, numbs2)

	var number = [] int {1, 2, 3, 4, 5}
	fmt.Println(number[0:2])
	fmt.Println(number[:2])
	fmt.Println(number[1:])

	var number1 = [] int {1, 2, 3}
	var number2 = [] int {1, 2, 3}

	fmt.Println(slices.Equal(number1, number2))

	var num_two_d = [][] int {{1, 2, 3}, {4, 5, 6}}
	fmt.Println(num_two_d)
}