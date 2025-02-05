package main

import "fmt"

func main() {
	jagged_array := [][]int{{1, 2, 3}, {1, 2}, {1, 2, 3, 4, 5, 6}}
	fmt.Println(jagged_array)

	// 1) Adding a new row with different size
	jagged_array = append(jagged_array, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	for i, row := range jagged_array {
		fmt.Printf("row %d : %v\n", i, row)
	}

	// 2) Adding an Element to a Specific Row
	row_index := 1
	jagged_array[row_index] = append(jagged_array[row_index], 10)
	fmt.Println(jagged_array)

	// 3) To remove an element from a specific row, we use slice manipulation.
	element_index := 1
	jagged_array[row_index] = append(jagged_array[row_index][:element_index], jagged_array[row_index][element_index+1:]...)
	fmt.Println(jagged_array)

	// 4) We iterate through all rows and columns dynamically.
	for i, row := range jagged_array {
		for j, val := range row {
			fmt.Printf("the row %d col %d value %v", i, j, val)
			fmt.Println()
		}
	}
}
