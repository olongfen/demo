package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	arr := []int{}
	cur := l
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return fmt.Sprintf("%v", arr)
}

func MakeList(val []int) *ListNode {
	if len(val) == 0 {
		return nil
	}
	root := &ListNode{
		Val:  val[0],
		Next: nil,
	}
	node := root
	for i := 1; i < len(val); i++ {
		node.Next = &ListNode{
			Val:  val[i],
			Next: nil,
		}
		node = node.Next
	}
	return root
}
