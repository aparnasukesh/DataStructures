package main

import "fmt"

func binarySearchRecursive(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearchRecursive(arr, target, low, mid-1)
	} else {
		return binarySearchRecursive(arr, target, mid+1, high)
	}
}

func binarySearch(arr []int, target int) int {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func main() {

	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	result := binarySearch(sortedArray, target)

	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
