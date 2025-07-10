package main

import "fmt"

func main() {
	//s := "abbaca"
	// Output: "ca"
	// Explanation:
	// For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
	// Example 2:

	s := "azxxzy"
	// Output: "ay"
	fmt.Println(removeDuplicates(s))

}

func removeDuplicates(s string) string {
	duplicate := true

	for len(s) > 2 && duplicate {
		duplicate = false
		for i := 0; i < len(s); i++ {
			if s[i] == s[i+1] {
				s = s[:i] + s[i+2:]
				duplicate = true
				fmt.Println("s: ", s)
				break
			}
		}
	}
	return s
}
