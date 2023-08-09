package main

import "fmt"

// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: true
// Example 2:

// Input: nums = [1,2,3,4]
// Output: false

func containsDuplicate(nums []int) bool {
	numberMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := numberMap[nums[i]]; ok {
			return true
		}
		numberMap[nums[i]] += 1
	}
	return false
}

func main() {

	nums := []int{1,2,3,1}
	fmt.Println(containsDuplicate(nums))
	
}