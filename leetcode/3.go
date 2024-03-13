package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("au"))
}

func lengthOfLongestSubstring(s string) int {
	var (
		left, right = 0, 0
		count       int
		maxCount    int
		existMap    = make(map[byte]int)
	)
	fc := func(count1, count2 int) int {
		if count1 >= count2 {
			return count1
		}
		return count2
	}
	for left <= right && right < len(s) {
		next := s[right]
		_, ok := existMap[next]
		if ok {
			left++
			right = left
			existMap = map[byte]int{}
			maxCount = fc(maxCount, count)
			count = 0
			continue
		}
		existMap[next] = 0
		count++
		right++
	}
	maxCount = fc(maxCount, count)
	return maxCount
}
