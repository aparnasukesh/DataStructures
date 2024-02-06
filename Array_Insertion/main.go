package main

import "fmt"

func main() {
	var pos, data int
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Enter the position")
	fmt.Scanf("%d", &pos)
	fmt.Println("Enter the value you want to insert")
	fmt.Scanf("%d", &data)

	if pos > 0 && pos <= len(arr)+1 {
		arr = append(arr, 0)
		for i := len(arr) - 1; i > pos-1; i-- {
			arr[i] = arr[i-1]
		}
		arr[pos-1] = data
		fmt.Println(arr)
	} else {
		fmt.Println("Invalid position")
	}

}
