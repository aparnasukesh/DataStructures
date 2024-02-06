package main

import "fmt"

func main() {
	var pos int
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Enter the position")
	fmt.Scanf("%d", &pos)

	if pos > 0 && pos <= len(arr) {
		for i := pos - 1; i < len(arr)-1; i++ {
			arr[i] = arr[i+1]
		}
		arr = arr[:len(arr)-1]
		fmt.Println(arr)
	} else {
		fmt.Println("Invalid position")
	}
}
