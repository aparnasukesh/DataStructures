package main

import "fmt"

func main() {
	var data int
	arr := []int{16, 2, 35, 90, 10, 3}
	fmt.Println("Enter the data you want to search")
	fmt.Scanf("%d", &data)
	flag := 0
	position := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == data {
			flag = 1
			position = i + 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("Data Not Found")
	} else {
		fmt.Printf("Data found at position %d", position)
	}
}
