package main

import "fmt"

func main() {
	arr := []int{7, 6, 10, 5, 9, 2, 1, 15, 7}
	fmt.Println("Array:", arr)
	res := quickSort(arr)
	fmt.Println("Sorted Array:", res)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[1]
	arr[len(arr)-1], arr[1] = arr[1], arr[len(arr)-1]

	var left, right []int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}
