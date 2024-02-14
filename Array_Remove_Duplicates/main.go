package main

import "fmt"

func main() {
	originalArray := []int{1, 2, 2, 3, 4, 4, 5, 6, 6, 7}
	uniqueArray := removeDuplicates(originalArray)

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Array without Duplicates:", uniqueArray)
}

func removeDuplicates(arr []int) []int {
	temp := make(map[int]bool)
	res := []int{}
	for _, val := range arr {
		if _, ok := temp[val]; !ok {
			temp[val] = true
			res = append(res, val)
		}
	}
	return res
}
