package main

import "fmt"

func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printGenericSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printCertainTypeSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printComparableSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printMultipleComparableSlice[T comparable, v string](items []T, name v) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack struct {
	elements []int
}

type stackAny [T any] struct {
	elements []T
}

func main() {
	printSlice([]int{1, 2, 3})
	printStringSlice([]string{"golang", "javascript"})
	printGenericSlice([]any{10, "typescript", 13.3, true})
	printCertainTypeSlice([]string{"Hello", "world"})

	myStack := stack{
		elements: []int{1, 2, 3},
	}

	myAnyStack := stackAny[any]{
		elements: []any {1, 2},
	}

	fmt.Println(myStack)
	fmt.Println(myAnyStack)
}
