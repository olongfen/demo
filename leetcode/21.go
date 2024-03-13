package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curl1 := list1
	curl2 := list2
	anwer := dummy
	for curl1 != nil || curl2 != nil {
		if curl1 != nil && curl2 != nil {
			if curl1.Val <= curl2.Val {
				anwer.Next = curl1
				curl1 = curl1.Next
			} else {
				anwer.Next = curl2
				curl2 = curl2.Next
			}
			anwer = anwer.Next
		} else if curl1 == nil && curl2 != nil {
			anwer.Next = curl2
			break
		} else {
			anwer.Next = curl1
			break
		}
	}
	return dummy.Next
}
