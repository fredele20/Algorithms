package solution

// This solution works perfect.
func Solution(N int) int {
	// Implement your solution here
	// This array is going to store binary digit
	var arr []int
	// Count holds the actual gap.
	var count int
	// Counter increments every time there is a zero
	var counter int
	// CurrentCount keep tracks of the current count
	var currentCount int

	// here we populate the array with the binary equivalent of the given number
	for N > 0 {
		arr = append([]int{N % 2}, arr...)
		N /= 2
	}

	// Here, we loop through the slice arr
	for i := 1; i < len(arr); i++ {
		// Check if the value of the current index is 0 and increment counter
		// also stores counter as the current count
		if arr[i] == 0 {
			counter++
			currentCount = counter
		}
		// Checks if the current count is greater than the count and
		// the current index value is 1, then stores currentCount to count
		if currentCount > count && arr[i] == 1 {
			count = currentCount
		}
		// Checks if the current index value is 1 and reset couter to zero
		// This helps to start counting from the next zero when the current value is 1
		if arr[i] == 1 {
			counter = 0
		}
	}
	return count
}
