package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}

	sum := 0

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	anoNums := []int{4, 5, 6}

	for _, num := range anoNums {
		sum = sum + num
		fmt.Println(sum)
	}

	num2 := []int{1, 2, 3}

	for i, num := range num2 {
		fmt.Println(num, i)
	}

	m := map[string]string{"fname": "ayush", "lname": "kumar"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for i, c := range "golang" {	// it will give the index and the unicode of the character
		fmt.Println(i, c)
	}

	// unicode -> 300 = 1 byte, 2 byte
}
