package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	data := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		var (
			start = i + 1
			end   = len(nums) - 1
		)
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return sum
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(data-target)) {
				data = sum
			}

		}
	}
	return data
}
