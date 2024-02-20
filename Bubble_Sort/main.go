package main

import "fmt"

func main() {
	arr := []int{15, 16, 6, 8, 5}
	fmt.Println(bubble_Sort(arr))
}

func bubble_Sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return arr
}
