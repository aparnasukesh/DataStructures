package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Enter a positive integer")
		return
	}

	first := 0
	second := 1

	fmt.Println("Fibonacci Series:")
	for i := 1; i <= n; i++ {
		fmt.Println(first)
		next := first + second
		first = second
		second = next
	}

}
