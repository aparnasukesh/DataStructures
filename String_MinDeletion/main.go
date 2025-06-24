package main

import (
	"fmt"
	"math"
)

func main() {

	// s := "abc"
	// k := 2

	// // Output: 1

	// // Explanation:

	// // s has three distinct characters: 'a', 'b' and 'c', each with a frequency of 1.
	// // Since we can have at most k = 2 distinct characters, remove all occurrences of any one character from the string.
	// // For example, removing all occurrences of 'c' results in at most k distinct characters. Thus, the answer is 1.
	// fmt.Println(minDeletion(s, k))
	s := "aabb"
	k := 2

	// Output: 0

	// Explanation:

	// s has two distinct characters ('a' and 'b') with frequencies of 2 and 2, respectively.
	// Since we can have at most k = 2 distinct characters, no deletions are required. Thus, the answer is 0.
	fmt.Println(minDeletion(s, k))

}

func minDeletion(s string, k int) int {
	char := make(map[rune]int)
	for _, ch := range s {
		char[ch]++
	}
	if len(char) <= k {
		return 0
	}
	count := len(char) - k
	ops := 0
	for count > 0 {
		small := math.MaxInt16
		var k rune
		for key, val := range char {
			if val <= small {
				small = val
				k = key
			}
		}
		delete(char, k)
		ops = ops + small
		count--
	}
	return ops
}
