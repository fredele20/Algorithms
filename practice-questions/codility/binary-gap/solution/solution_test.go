package solution

import "testing"

type TestData struct {
	num int
	gap int
}

func TestSolution(t *testing.T) {

	tests := []TestData{
		{num: 1376796946, gap: 5},
		{num: 1041, gap: 5},
		{num: 32, gap: 0},
		{num: 15, gap: 0},
		{num: 42, gap: 1},
		{num: 66561, gap: 9},
		{num: 74901729, gap: 4},
		{num: 805306373, gap: 25},
		{num: 1376796946, gap: 5},
	}

	for _, test := range tests {
		gapResult := Solution(test.num)
		if gapResult != test.gap {
			t.Errorf("Solution(%d) FAILED, Expected %d, got %d", test.num, test.gap, gapResult)
		}
	}

}
