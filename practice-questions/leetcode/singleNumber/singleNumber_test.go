package main

import "testing"

type testData struct {
	nums []int
	num  int
}

func TestSingleNumber(t *testing.T) {
	tests := []testData{
		{nums: []int{4, 2, 1, 2, 1}, num: 4},
		{nums: []int{3, 1, 5, 2, 1, 0, 3, 1, 5}, num: 0},
		{nums: []int{5, 2, 3, 1, 3, 1, 5}, num: 2},
		{nums: []int{3, 9, 3, 4, 4}, num: 9},
	}

	for _, test := range tests {
		result := SingleNumber(test.nums)
		if result != test.num {
			t.Errorf("SingleNumber(%v) FAILED, Expected %d, got %d\n", test.nums, test.num, result)
		}
	}
}
