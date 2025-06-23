package main

import (
	"fmt"
	"strings"
)

func main() {
	caption := "Leetcode daily streak achieved"

	// Output: "#leetcodeDailyStreakAchieved"

	// Explanation:

	// The first letter for all words except "leetcode" should be capitalized.

	//caption := "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"

	// Output: "#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"

	// Explanation:

	// Since the first word has length 101, we need to truncate the last two letters from the word.
	//caption := "Bold apple beyond bright future crash mountains glow light gently dance waits shore breeze mind "
	fmt.Println(generateTag(caption))
}

func generateTag(caption string) string {
	words := strings.Split(caption, " ")
	fmt.Println(words)
	str := "#"
	for i := 0; i < len(words); i++ {
		if len(words[i]) != 0 {
			if i == 0 || len(str) == 1 {
				ch1 := strings.ToLower(string(words[i][0]))
				ch2 := strings.ToLower(string(words[i][1:]))
				word := ch1 + ch2
				str = str + word
			} else {
				ch1 := strings.ToUpper(string(words[i][0]))
				ch2 := strings.ToLower(string(words[i][1:]))
				word := ch1 + ch2
				str = str + word
			}
		}
	}
	if len(str) > 100 {
		diff := len(str) - 100
		str = str[:len(str)-diff]
	}
	return str
}
