package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	node := &ListNode{
		Val:  0,
		Next: head,
	}
	second := node
	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return node.Next
}
