// package main

// import "fmt"

// func main() {
// 	arr := []int{7, 6, 10, 5, 9, 2, 1, 15, 7}
// 	low := 0
// 	high := len(arr) - 1
// 	quickSort(arr, low, high)
// 	fmt.Println(arr)

// }

// func quickSort(arr []int, low, high int) {

// 	if low < high {
// 		pivotIndex := partition(arr, low, high)
// 		quickSort(arr, low, pivotIndex-1)
// 		quickSort(arr, pivotIndex+1, high)
// 	}

// }

// func partition(arr []int, low, high int) int {
// 	pivot := arr[high]
// 	i := low - 1

// 	for j := low; j < high; j++ {
// 		if arr[j] <= pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}

// 	arr[i+1], arr[high] = arr[high], arr[i+1]

// 	return i + 1
// }

// // func partition(arr []int, low, high int) int {
// // 	pivot := arr[low]
// // 	i := low + 1
// // 	j := high

// // 	for {
// // 		for i <= j && arr[i] <= pivot {
// // 			i++
// // 		}

// // 		for i <= j && arr[j] > pivot {
// // 			j--
// // 		}

// // 		if i <= j {
// // 			arr[i], arr[j] = arr[j], arr[i]
// // 		} else {
// // 			break
// // 		}
// // 	}

// // 	arr[low], arr[j] = arr[j], arr[low]

// // 	return j
// // }

package main

import "fmt"

func Quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(Quicksort(left), pivot), Quicksort(right)...)
}

func main() {
	arr := []int{7, 6, 10, 5, 9, 2, 1, 15, 7}
	result := Quicksort(arr)
	fmt.Println(result)
}
