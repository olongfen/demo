package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	arr := [][]int{}
	l := len(nums)
	if l < 3 {
		return arr
	}
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			return arr
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := l - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				arr = append(arr, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}

		}
	}
	return arr
}
