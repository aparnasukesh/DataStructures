package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aassddhhakjaaaahhaa"
	subStr := "aa"
	count := strings.Count(str, subStr)
	fmt.Println(count)

	fmt.Println(CountSubstring(str, subStr))
}

func CountSubstring(str, substring string) int {
	count := 0

	for i := 0; i <= len(str)-len(substring); i++ {
		if str[i:i+len(substring)] == substring {
			count++
		}
	}
	return count
}
