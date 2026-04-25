package main

import "fmt"

func main() {
	// add 0 as default value as it is int
	var nums [4]int

	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(nums)

	// add false as default value as it is int
	var vals [4] bool
	fmt.Println(vals)

	// give array with empty string
	var str [4] string
	fmt.Println(str)

	str[2] = "golang"
	fmt.Println(str)

	// to declare in single line
	numbers := [4] int {1, 2, 3, 4}
	fmt.Println(numbers)

	// 2d arrays
	num_two_dimension := [2][2] int {{1, 2}, {3, 4}}
	fmt.Println(num_two_dimension)

	// fixed size, that is predictable
	// Memory optimization
	// constant time access
}