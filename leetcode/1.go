package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1}, 11))
}

func twoSum(nums []int, target int) []int {
	var (
		m = make(map[int]int)
	)
	for i, v := range nums {
		val, ok := m[target-v]
		if ok && nums[val]+v == target {
			return []int{i, val}
		}
		m[target-v] = i
	}
	ret := []int{}
	for i, v := range nums {
		val, ok := m[v]
		if ok && val != i {

			ret = append(ret, val)

		}
	}
	return ret
}
