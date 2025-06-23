package main

import "fmt"

func main() {
	arr := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	// Output: true
	// Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
	// Example 2:

	// Input: arr = [0,2,1,-6,6,7,9,-1,2,0,1]
	// Output: false
	// Example 3:

	// Input: arr = [3,3,6,5,-2,2,5,1,-9,4]
	// Output: true
	// Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
	fmt.Println(canThreePartsEqualSum(arr))

}

func canThreePartsEqualSum(arr []int) bool {
	totalSum := sumArr(arr)
	if totalSum%3 != 0 {
		return false
	}
	target := totalSum / 3
	first := 0
	i := 0
	for _, a := range arr[:len(arr)-1] {
		first += a
		if first == target {
			second := 0
			j := i + 1
			for _, b := range arr[i+1 : len(arr)-1] {
				second += b
				if second == target {
					third := sumArr(arr[j+1:])
					if third == target {
						return true
					}
				}
				j++
			}
		}
		i++
	}

	return false
}

func sumArr(arr []int) int {
	totalSum := 0
	for i := 0; i < len(arr); i++ {
		totalSum = totalSum + arr[i]
	}
	return totalSum
}
