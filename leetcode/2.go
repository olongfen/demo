package main

import (
	"fmt"
)

func main() {
	l1 := MakeList([]int{2, 4, 9})
	l2 := MakeList([]int{5, 6, 4, 9})
	l := addTwoNumbers(l1, l2)
	fmt.Println(l)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位
	carry := 0
	indexL1 := l1
	indexL2 := l2

	for indexL1 != nil && indexL2 != nil {
		sum := indexL1.Val + indexL2.Val + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		indexL1.Val = sum % 10
		if indexL1.Next == nil && indexL2.Next != nil {
			indexL1.Next = indexL2.Next
			indexL1 = indexL1.Next
			break
		}
		if indexL1.Next == nil && indexL2.Next == nil && carry == 1 {
			indexL1.Next = &ListNode{
				Val:  carry,
				Next: nil,
			}
			carry = 0
		}
		indexL1 = indexL1.Next
		indexL2 = indexL2.Next
	}
	for indexL1 != nil && carry > 0 {
		sum := indexL1.Val + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		indexL1.Val = sum % 10
		if indexL1.Next == nil && carry > 0 {
			indexL1.Next = &ListNode{
				Val: carry,
			}
			carry = 0
			break
		}
		indexL1 = indexL1.Next
	}
	if carry == 1 {
		indexL1.Next = &ListNode{
			Val: carry,
		}
	}
	return l1
}
