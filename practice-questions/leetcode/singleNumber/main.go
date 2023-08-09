package main

import "fmt"

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

// Example:
// Input: nums = [4,1,2,1,2]
// Output: 4
func singleNumber(nums []int) int {
  var singleNumber int
  numberMap := make(map[int]int)

  for i := 0; i < len(nums); i++ {
    numberMap[nums[i]] += 1
  }
  for key, val := range numberMap {
    if val == 1 {
      singleNumber = key
    }
  }

  return singleNumber
}

func main() {

	arr := []int{4,1,2,1,2}
	fmt.Println(singleNumber(arr))
	
}