package main

import "fmt"

func main() {
	fmt.Println(func1(5))
}

func func1(n int) int {
	if n <= 1 {
		return 1
	}
	return n * func2(n-1)
}

func func2(n int) int {
	if n <= 1 {
		return 1
	}
	return n * func1(n-1)
}
