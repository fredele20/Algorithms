package main

import "fmt"

func cyclicRotation(nums []int, k int) []int {
	copySlice := []int{}
	newSlice := []int{}
	lastIndex := len(nums) -1
	if k == 0 {
		return nums
	}
	for i := 1; i <= k; i++ {
		for j := 0; j <= lastIndex; j++ {
			if j != lastIndex {
				copySlice = append(copySlice, nums[j])
			}
			newSlice = append([]int{nums[j]}, copySlice...)
		}
		copySlice = []int{}
		nums = newSlice
	}
	return newSlice
}

func main() {
	nums := []int{1, 2, -3, 4}
	k := 4
	fmt.Println(cyclicRotation(nums, k))
}