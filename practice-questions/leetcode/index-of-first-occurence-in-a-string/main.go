package main

import (
	"fmt"
	"strings"
)

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, 
// or -1 if needle is not part of haystack.

// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {

	str1 := "sadbutsad"
	str2 := "sad"
	fmt.Println(strStr(str1, str2))
}
