package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right >= left {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	nums := []int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50}
	fmt.Println(len(nums))
	target := 31

	output := searchInsert(nums, target)
	fmt.Println("Output: ", output)

	x := 1 /2
	fmt.Println("x: ", x)
}
