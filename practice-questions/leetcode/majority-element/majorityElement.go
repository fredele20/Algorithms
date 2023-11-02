package main


func MajorityElement(arr []int) int {

	numberMap := make(map[int]int)
	majority := 0
	max := 0

	if len(arr) == 1 {
		max = arr[0]
		return max
	}

	for _, num := range arr {
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