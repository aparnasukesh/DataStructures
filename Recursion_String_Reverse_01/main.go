package main

import "fmt"

func main() {
	s := "aparna"
	fmt.Println(Reverse(s))
}

func Reverse(input string) string {
	if len(input) <= 1 {
		return input
	}

	return Reverse(input[1:]) + string(input[0])
}
