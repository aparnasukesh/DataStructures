package main

import "fmt"

func main() {
	//a := []int{1, 2, 3, 4, 5, 6}
	var size int

	fmt.Println("Enter the size of the array")
	fmt.Scanf("%d", &size)
	a := make([]int, size)
	fmt.Println("Eneter the values")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Printf("%d", a[0])
	fmt.Println("The Array:")
	for _, val := range a {
		fmt.Println(val)
	}
}
