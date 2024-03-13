package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	if len(str) == 0 {
		return false
	}
	l := 0
	r := len(str) - 1
	for l < r {
		if str[l] == str[r] {
			l++
			r--
		} else {
			return false
		}

	}
	return true
}
