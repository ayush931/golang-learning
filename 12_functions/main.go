package main

import "fmt"

// if type of both the parameter is same
func add(a, b int) int {
	return a + b
}

// multiple return value
func getLanguages() (string, string, string, bool) {
	return "golang", "javascript", "java", false
}

func processInt(fn func(a int) int) {
	fn(1)
}

func processIt() func(a int) int {
	return  func(a int) int {
		return 2
	}
} 

func main() {
	ans := add(4, 5)
	fmt.Println(ans)
	fmt.Println(getLanguages())

	// taking multiple return value in multiple variables
	lang1, lang2, lang3, _ := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	fn := func (a int) int {
		return ans
	}

	processInt(fn)
	f1 := processIt()
	fmt.Println(f1(2))
}