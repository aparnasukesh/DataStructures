package main

import "fmt"

func main() {
	fmt.Println(print(100))
}

func print(a int) int {
	if a < 1 {
		return 1
	}
	fmt.Println("a: ", a)
	return a + print(a/2)
}
