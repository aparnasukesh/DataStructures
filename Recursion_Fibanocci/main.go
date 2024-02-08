package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {

	numTerms := 10

	fmt.Printf("Fibonacci series up to %d terms:\n", numTerms)
	for i := 0; i < numTerms; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
