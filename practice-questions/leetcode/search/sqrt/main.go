package main

import "fmt"

func mySqrt(x int) int {
  left, right := 0, x + 1

  for right > left {
    mid := left + (right - left) / 2

    if mid * mid > x {
      right = mid
    } else {
      left = mid + 1
    }
  }
  return left - 1
}

func plusOne(digits []int) []int {
  lastIndex := len(digits) - 1

  if digits[lastIndex] != 9 {
    digits[lastIndex]++
    return digits
  } else {
    digits[lastIndex] = 0
    digits = append([]int{1}, digits...)
    return digits
  }
}

func main() {
	x := mySqrt(5)
	nums := []int{1, 9}

	output := plusOne(nums)
	fmt.Println("Plus one output: ", output)

	fmt.Println("Output: ", x)
}