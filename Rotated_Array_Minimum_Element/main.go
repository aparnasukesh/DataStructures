package main

func main() {
	// Input: nums = [3,4,5,1,2]
	// Output: 1
	// Explanation: The original array was [1,2,3,4,5] rotated 3 times.
	// Example 2:

	// Input: nums = [4,5,6,7,0,1,2]
	// Output: 0
	// Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
	// Example 3:

	// Input: nums = [11,13,15,17]
	// Output: 11
	// Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

}

// Find minimum element from a rotated sorted list
func findMin(nums []int) int {
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[1]
		}
		return nums[0]
	}
	l, r := 0, len(nums)-1
	min := nums[0]
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < min {
			min = nums[mid]
		} else if nums[mid] > min {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return min
}
