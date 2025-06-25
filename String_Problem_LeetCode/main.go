package main

import (
	"fmt"
)

func main() {

	//Example 1:

	columnTitle := "A"
	// Output: 1
	// Example 2:

	//columnTitle := "AB"
	// // Output: 28
	// // Example 3:

	//columnTitle := "ZY"
	// //Output: 701
	fmt.Println(titleToNumber(columnTitle))
	fmt.Println(convertToTitle(28))

}

func titleToNumber(columnTitle string) int {

	multiplier := 1
	result := 0

	for i := len(columnTitle) - 1; i >= 0; i-- {
		result = result + (int(columnTitle[i])-int('A')+1)*multiplier
		multiplier = multiplier * 26
	}

	return result
}

func convertToTitle(columnNumber int) string {
	finalStr := ""

	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		finalStr = string('A'+remainder) + finalStr
		fmt.Println(finalStr)
		columnNumber /= 26
	}

	return finalStr
}
