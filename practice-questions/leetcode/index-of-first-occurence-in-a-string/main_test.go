package main

import "testing"

type testData struct {
	str1  string
	str2  string
	index int
}

func TestStrstr(t *testing.T) {
	tests := []testData{
		{str1: "sadbutsad", str2: "sad", index: 0},
		{str1: "bigdaddy", str2: "dad", index: 3},
		{str1: "sadbutsad", str2: "baddy", index: -1},
		{str1: "alloverlove", str2: "love", index: 2},
		{str1: "wrothenfriedegg", str2: "hen", index: 4},
		{str1: "allthebest", str2: "", index: -1},
		{str1: "", str2: "", index: -1},
	}

	for _, test := range tests {
		result := StrStr(test.str1, test.str2)
		if result != test.index {
			t.Errorf("Strstr(%v, %v) FAILED! expected %d, got %d", test.str1, test.str2, test.index, result)
		}
	}
}
