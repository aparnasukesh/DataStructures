package main

import (
	"fmt"
	"sort"
)

func main() {
	var data int
	var mid int

	arr := []int{16, 2, 35, 90, 10, 3, 8}
	sort.Ints(arr)
	fmt.Println(arr)

	fmt.Println("Enter the data you want to search")
	fmt.Scanf("%d", &data)

	l := 0
	r := len(arr) - 1
	flag := 0
	for l <= r {
		mid = (l + r) / 2
		if data == arr[mid] {
			flag = 1
			break
		} else if data < arr[mid] {
			r = mid - 1
		} else if data > arr[mid] {
			l = mid + 1
		}
	}

	if flag == 0 {
		fmt.Println("Data not found")
	} else {
		fmt.Printf("Data found at position %d", mid)
	}
}
