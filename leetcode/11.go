package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 2, 4, 3}))
}

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1
	maxVal := 0
	for left < right {
		val := (right - left) * min(height[left], height[right])
		maxVal = max(maxVal, val)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
