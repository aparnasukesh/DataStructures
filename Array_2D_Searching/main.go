package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	//Output: true
	// matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	// target := 13
	//Output: false
	fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	i := 0
	for i < len(matrix) {
		if target == matrix[i][len(matrix[0])-1] {
			return true
		} else if target < matrix[i][len(matrix[0])-1] {
			for j := 0; j < len(matrix[i]); j++ {
				if target == matrix[i][j] {
					return true
				}
			}
		}
		i++
	}
	return false
}
