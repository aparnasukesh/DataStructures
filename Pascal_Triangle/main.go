package main

import "fmt"

func main() {
	numRows := 5
	// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	// Example 2:

	//numRows := 1
	//Output: [[1]]
	fmt.Println(generate(numRows))

}
func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				temp[j] = 1
			} else {

				temp[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res = append(res, temp)
	}
	return res
}
