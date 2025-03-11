package main

import "fmt"

func main() {
	fmt.Println(sum(2))
}

func sum(n int) int {
	if n <= 1 {
		return n
	}
	return n + sum(n-1)
}
