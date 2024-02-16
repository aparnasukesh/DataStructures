package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 2, 6, 9, 8}
	large := arr[0]
	secondLarge := arr[1]

	if secondLarge > large {
		large, secondLarge = secondLarge, large
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] > large {
			secondLarge = large
			large = arr[i]
		} else if arr[i] > secondLarge && arr[i] < large {

			secondLarge = arr[i]
		}
	}
	fmt.Println(secondLarge)
}
