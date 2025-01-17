package main

import "fmt"

func main() {
	//s := "()"
	s := "()[]{}"
	fmt.Println(isValid(s))
	//Output: true
}

func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, char := range s {
		fmt.Println(stack)
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
			fmt.Println(stack)

		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(stack)

			if bracketMap[char] != top {
				return false
			}
		}
	}
	return len(stack) == 0
}
