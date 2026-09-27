package main

import "fmt"

// taking unlimited parameter of given type
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

// func anotherSum(nums ...interface{}) int { -> unlimited argument of any type

func main() {
	// unlimited parameters
	fmt.Println(1, 2, 3, 4, 5)

	result := sum(1, 2, 3, 4, 5, 6)

	nums := []int{5, 6, 7, 8, 9}
	results := sum(nums...)

	fmt.Println(result)
	fmt.Println(results)
}
