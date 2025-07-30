package main

func main() {
	// Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
	// Output: true

}
func checkStraightLine(coordinates [][]int) bool {
	x0, x1 := coordinates[0][0], coordinates[1][0]
	y0, y1 := coordinates[0][1], coordinates[1][1]

	for i := 1; i < len(coordinates); i++ {
		x2, y2 := coordinates[i][0], coordinates[i][1]
		if (x1-x0)*(y2-y1) != (y1-y0)*(x2-x1) {
			return false
		}
	}
	return true
}
