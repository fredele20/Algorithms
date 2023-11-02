package main

import "testing"

type testData struct {
	arr []int
	output int
}

func TestMajorityElement(t *testing.T) {

	tests := []testData{
		{arr: []int{1, 2, 1, 2, 1}, output: 1},
		{arr: []int{3, 1, 5, 2, 1, 0, 3, 1, 5}, output: 1},
		{arr: []int{5, 2, 3, 1, 3, 1, 5, 3}, output: 3},
		{arr: []int{3, 4, 9, 3, 4, 4}, output: 4},
	}

	for _, test := range tests {
		result := MajorityElement(test.arr)
		if result != test.output {
			t.Errorf("MajorityELement(%v) FAILED, expected %d, got %d", test.arr, test.output, result)
		}
	}

}
