package main

import "fmt"

// 反转从位置 left 到位置 right 的链表节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	// 创建一个虚拟头节点，以处理 left == 1 的情况
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	prev := dummy

	// 找到反转的起始节点前一个节点 prev
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// 初始化反转区间的两个指针
	current := prev.Next
	next := current.Next

	// 开始反转
	for i := 0; i < right-left; i++ {
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = current.Next
	}

	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	fmt.Println(reverseBetween(head, 2, 4))
}
