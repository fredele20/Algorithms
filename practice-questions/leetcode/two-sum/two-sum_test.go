package main

import (
	"reflect"
	"testing"
)

type testData struct {
	nums   []int
	target int
	output []int
}

func TestTwoSum(t *testing.T) {
	tests := []testData{
		{nums: []int{2, 7, 11, 15}, target: 9, output: []int{0, 1}},
		{nums: []int{2, 5, 13, 1}, target: 15, output: []int{0, 2}},
		{nums: []int{3, 6, 8, 4}, target: 10, output: []int{1, 3}},
		{nums: []int{3, 6, 8, 4}, target: 11, output: []int{0, 2}},
		{nums: []int{3, 6, 8, 4}, target: 15, output: []int{}},
		{nums: []int{}, target: 11, output: []int{}},
		{nums: []int{3, 6, 8, 4}, output: []int{}},
	}

	for _, test := range tests {
		result := TwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("twoSum(%v, %d) FAILED, expected %v, got %v", test.nums, test.target, test.output, result)
		}

		for i, v := range result {
			if v != test.output[i] {
				t.Errorf("twoSum(%v, %d) FAILED, expected %d, got %v", test.nums, test.target, v, test.output[i])
			}
		}
	}
}
