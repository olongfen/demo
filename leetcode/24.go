package main

import "fmt"

func main() {
	l1 := MakeList([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(swapPairs(l1))
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		next := head.Next.Next
		pre.Next = head.Next
		// 防止变成环
		head.Next = next
		pre.Next.Next = head
		pre = pre.Next.Next
		head = next
	}
	return dummy.Next
}
