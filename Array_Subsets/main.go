package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println(result)
}

func subsets(nums []int) [][]int {
	result := [][]int{{}} // start with an empty subset

	for _, num := range nums {
		newSubsets := [][]int{}
		for _, subset := range result {
			// Create a new subset that includes the current number
			newSubset := append([]int{}, subset...) // copy existing subset
			newSubset = append(newSubset, num)
			fmt.Println("\nnewsubset: ", newSubset)
			newSubsets = append(newSubsets, newSubset)
		}
		// Add all new subsets to result
		fmt.Println("\nnewsubsets: ", newSubsets)
		result = append(result, newSubsets...)
		fmt.Println("\nresult:", result)
	}

	return result
}
