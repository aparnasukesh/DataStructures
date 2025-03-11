package main

import "fmt"

func main() {
	display(5)
}

func display(n int) {
	if n < 1 {
		return
	} else {
		fmt.Printf("%d", n)
		display(n - 1)
		fmt.Printf("%d", n)
	}
}
