package main

import "fmt"

func main() {
	s := "abab"
	// Output: true
	// Explanation: It is the substring "ab" twice.
	fmt.Println(repeatedSubstringPattern(s))
}

func repeatedSubstringPattern(s string) bool {
	ss := s
	for i := 0; i < len(s)-1; i++ {
		s = string(s[len(s)-1]) + string(s[:len(s)-1])
		if ss == s {
			return true
		}
	}

	return false
}
