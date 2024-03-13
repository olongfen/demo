package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "fle", "flood"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, i := range strs {
		// 循环遍历是否存在相同的前缀
		for strings.Index(i, prefix) != 0 {
			// 遍历完没有相同前缀，直接return
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix

}
