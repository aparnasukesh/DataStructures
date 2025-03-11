package main

import "fmt"

func main() {

	print(100)
}

func print(a int) {
	if a < 1 {
		return
	}
	fmt.Println(a)
	print(a / 2)
}
