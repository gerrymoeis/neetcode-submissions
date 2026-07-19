func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var head *ListNode
	small := list1
	big := list2
	if small.Val >= big.Val {
		head = big
	} else {
		head = small
	}
	for small != nil && big != nil {
		if small.Val >= big.Val {
			small, big = big, small
		}
		nextNode := small.Next
		if nextNode == nil || nextNode.Val >= big.Val {
			small.Next = big
		}
		small = nextNode
	}
	return head
}