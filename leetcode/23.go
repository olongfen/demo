package main

func mergeKLists(lists []*ListNode) *ListNode {

	return mergeK_(lists, 0, len(lists)-1)
}

func mergeK_(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists(mergeK_(lists, l, mid), mergeK_(lists, mid+1, r))
}
