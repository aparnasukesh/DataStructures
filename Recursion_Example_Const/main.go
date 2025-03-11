package main

import "fmt"

func main() {
	count(3)
}

func count(n int) {
	d := 1
	fmt.Println("n: ", n)
	fmt.Println("d: ", d)
	d++
	if n > 1 {
		count(n - 1)
	}
	fmt.Println("d: ", d)
}
