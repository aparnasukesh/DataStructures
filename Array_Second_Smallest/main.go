package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 2, 6, 9, 8}

	// Initialize small and secondSmall with the first two elements
	small, secondSmall := arr[0], arr[1]

	// Swap if necessary to ensure small is smaller than secondSmall
	if secondSmall > small {
		secondSmall, small = small, secondSmall
	}

	// Iterate through the array starting from the third element
	for i := 2; i < len(arr); i++ {
		// If the current element is smaller than the current smallest
		if arr[i] < small {
			// Update secondSmall and small
			secondSmall = small
			small = arr[i]
		} else if arr[i] < secondSmall && arr[i] > small {
			// If the current element is smaller than secondSmall but greater than small
			// Update secondSmall
			secondSmall = arr[i]
		}
	}

	// Print the second smallest element
	fmt.Println(secondSmall)
}
