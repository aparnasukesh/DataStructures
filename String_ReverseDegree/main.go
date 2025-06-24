package main

import (
	"fmt"
	"unicode"
)

func main() {

	s := "abc"

	// Output:
	//
	//	148
	fmt.Println(reverseDegree(s))
}

func reverseDegree(s string) int {
	sum := 0
	j := 0

	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		fmt.Println("char: ", char)
		fmt.Println("unicode:", int(unicode.ToLower(char)), int('z'))

		position := int(unicode.ToLower(char)) - int('z') - 1
		fmt.Println("position: ", position)
		j = j + 1
		fmt.Println("j: ", j)
		sum = sum + abs(position)*j
		fmt.Println("sum: ", sum)

	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
