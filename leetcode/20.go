package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(isValid("()[]{}(]"))
}

func isValid(s string) bool {
	stack := list.New()
	ma := map[int32]int32{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack.PushFront(v)
		} else {
			ele := stack.Front()
			if ele == nil {
				return false
			}
			stack.Remove(ele)
			if ma[v] != ele.Value.(int32) {
				return false
			}
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}
