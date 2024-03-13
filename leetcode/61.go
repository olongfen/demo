package main

import "fmt"

func rotateRight(head *ListNode, k int) *ListNode {
	// 处理边界情况
	if head == nil || k == 0 {
		return head
	}

	// 计算链表的长度
	length := 1
	current := head
	for current.Next != nil {
		current = current.Next
		length++
	}

	// 计算实际需要向右移动的步数
	k %= length

	// 如果k为0，不需要移动
	if k == 0 {
		return head
	}

	// 找到新的头节点和新的尾节点
	newTail := head
	for i := 1; i < length-k; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	// 断开循环，将新的尾节点的Next指针设置为nil
	newTail.Next = nil
	printList(newHead)
	printList(newTail)
	// 将新的头节点连接到原始链表的头节点之前
	current = newHead
	for current.Next != nil {
		current = current.Next
	}
	current.Next = head
	return newHead
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// 创建一个链表 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	k := 2 // 向右移动的步数

	fmt.Println("原始链表:")
	printList(head)

	// 向右移动链表
	newHead := rotateRight(head, k)

	fmt.Println("向右移动后的链表:")
	printList(newHead)
}
