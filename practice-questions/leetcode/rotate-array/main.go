package main

import "fmt"

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		nums = append([]int{nums[len(nums) - 1]}, nums...)
		nums = nums[:len(nums)-1]
	}

	fmt.Println("Nums: ", nums)
}

func main() {
	nums := []int{-1,-100,3,99}
	k := 2
	rotate(nums, k)
}
