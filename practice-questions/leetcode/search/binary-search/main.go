package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	
	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1

}

func main() {

	nums := []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50}
	target := 40

	output := search(nums, target)
	fmt.Println("Output: ", output)

}
