package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcde"
	goal := "cdeab"
	// Output: true
	// Example 2:

	// Input: s = "abcde", goal = "abced"
	// Output: false
	fmt.Println(rotateString(s, goal))
}
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}
