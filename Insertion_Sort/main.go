package main

import "fmt"

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(arr)

}

func insertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && val < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
	fmt.Println(arr)
}
