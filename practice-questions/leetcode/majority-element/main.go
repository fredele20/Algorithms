package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	numberMap := make(map[int]int)
	var majority int
	var max int
	if len(nums) == 1 {
		majority = nums[0]
		return majority
	}
	for _, num := range nums {
		numberMap[num] += 1
	}

	for i, v := range numberMap {
		if v > majority {
			majority = v
			max = i
		}
	}
	return max
}

func main() {

	arr := []int{3,3,4}
	result := majorityElement(arr)
	fmt.Println(result)

}
