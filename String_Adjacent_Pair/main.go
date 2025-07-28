package main

import "fmt"

func main() {
	s := "leEeetcode"
	// Output: "leetcode"
	// Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
	// Example 2:
	s0 := "abBAcC"
	// Output: ""
	// Explanation: We have many possible scenarios, and all lead to the same answer. For example:
	// "abBAcC" --> "aAcC" --> "cC" --> ""
	// "abBAcC" --> "abBA" --> "aA" --> ""
	// Example 3:
	s1 := "s"
	//Output: "s"

	fmt.Println("example 1:", makeGood(s))
	fmt.Println("example 2: ", makeGood(s0))
	fmt.Println("example 3: ", makeGood(s1))

}
func makeGood(s string) string {
	stack := []rune{}

	for _, ch := range s {
		n := len(stack)
		if n > 0 && abs(int(ch)-int(stack[n-1])) == 32 {
			fmt.Println("stack: ", string(stack))
			stack = stack[:n-1]
		} else {
			stack = append(stack, ch)
			fmt.Println("stk: ", string(stack))
		}
	}
	return string(stack)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
