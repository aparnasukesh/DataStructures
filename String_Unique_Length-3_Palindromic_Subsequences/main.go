package main

import "fmt"

func main() {
	//Example 1:

	s := "aabca"
	// Output: 3
	// Explanation: The 3 palindromic subsequences of length 3 are:
	// - "aba" (subsequence of "aabca")
	// - "aaa" (subsequence of "aabca")
	// - "aca" (subsequence of "aabca")
	// Example 2:

	//s := "adc"
	// Output: 0
	// Explanation: There are no palindromic subsequences of length 3 in "adc".
	// Example 3:

	//s := "bbcbaba"
	// Output: 4
	// Explanation: The 4 palindromic subsequences of length 3 are:
	// - "bbb" (subsequence of "bbcbaba")
	// - "bcb" (subsequence of "bbcbaba")
	// - "bab" (subsequence of "bbcbaba")
	// - "aba" (subsequence of "bbcbaba")

	fmt.Println(countPalindromicSubsequence(s))

}

// func countPalindromicSubsequence(s string) int {
// 	check := make(map[string]bool)

// 	for i := 0; i < len(s)-2; i++ {
// 		j := len(s) - 1
// 		match := false

// 		for j >= i+2 && j >= 2 {
// 			if s[i] == s[j] {
// 				fmt.Println("s[i]:", string(s[i]), "s[j]: ", string(s[j]))
// 				match = true
// 				break
// 			}
// 			j--
// 		}
// 		if match {
// 			str := s[i : j+1]
// 			for k := 1; k < len(str)-1; k++ {
// 				sub := string(str[0]) + string(str[k]) + string(str[len(str)-1])
// 				if !check[sub] {
// 					check[sub] = true
// 				}
// 			}
// 		}
// 	}

// 	return len(check)

// }
func countPalindromicSubsequence(s string) int {
	fmt.Println("s: ", s)
	first := make([]int, 26)
	last := make([]int, 26)

	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}
	fmt.Println("first: ", first)
	fmt.Println("last: ", last)
	// Fill first and last occurrences
	for i := 0; i < len(s); i++ {
		fmt.Println("s[i]: ", s[i])
		c := int(s[i] - 'a')
		fmt.Println("c: ", c)
		if first[c] == -1 {
			first[c] = i
		}
		last[c] = i
	}
	fmt.Println("first: ", first)
	fmt.Println("last: ", last)
	result := 0

	// For each character a...z
	for c := 0; c < 26; c++ {
		if first[c] != -1 && last[c]-first[c] > 1 {
			// Find unique middle characters
			seen := make(map[byte]bool)
			for mid := first[c] + 1; mid < last[c]; mid++ {
				seen[s[mid]] = true
			}
			result += len(seen)
		}
	}

	return result
}
