package main

import "fmt"

// This solution has only 60% correctness, will work on it.
func solution(N int) int {
	// Implement your solution here
	var arr []int
	var count int
	var counter int
	var currentCount int

	for N != 0 {
		arr = append([]int{N % 2}, arr...)
		N /= 2
	}

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			counter++
			currentCount = counter
			if arr[len(arr)-1] == 0 {
				currentCount = 0
			}
			if currentCount > count {
				count = currentCount
			}
		}
		if arr[i] == 1 {
			counter = 0
		}
	}
	return count
}

func main() {

	n := solution(32)
	fmt.Println(n)

}
