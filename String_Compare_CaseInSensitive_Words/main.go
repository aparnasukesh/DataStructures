package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	//banned := []string{"hit"}
	paragraph := "a, a, a, a, b,b,b,c, c"
	banned := []string{"a"}
	//Output:"b"
	fmt.Println(mostCommonWord(paragraph, banned))
}

func mostCommonWord(paragraph string, banned []string) string {
	var builder strings.Builder
	for _, r := range paragraph {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(' ')
		}
	}
	paragraph = builder.String()
	check := make(map[string]bool)
	for _, w := range banned {
		check[strings.ToLower(w)] = true
	}
	words := strings.Fields(paragraph)
	count := 0
	ans := ""
	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		if !check[word] {
			count1 := 1
			for j := 0; j < len(words); j++ {
				if i != j && word == strings.ToLower(words[j]) {
					count1++
				}
			}
			if count < count1 {
				count = count1
				ans = word
			}
		}

	}
	return ans
}
