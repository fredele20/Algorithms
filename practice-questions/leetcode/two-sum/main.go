package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func twoSum(nums []int, target int) []int {
  var numMap = make(map[int]int)

  for i, num := range nums {
    if requiredNumIndex, isPresent := numMap[target - num]; isPresent {
      return []int{ requiredNumIndex, i }
    }
    numMap[num]=i
  }

  return []int{}
}

func main() {

	arr := []int{3,2,4}
	target := 6
	fmt.Println(twoSum(arr, target))
	
}