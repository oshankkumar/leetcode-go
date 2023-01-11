package merge_two_sorted_lists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		newHead *ListNode
	)

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	itr1, itr2 := list1, list2
	if itr1.Val < itr2.Val {
		newHead = &ListNode{Val: itr1.Val}
		itr1 = itr1.Next
	} else {
		newHead = &ListNode{Val: itr2.Val}
		itr2 = itr2.Next
	}

	itrNew := newHead
	for itr1 != nil && itr2 != nil {
		if itr1.Val < itr2.Val {
			itrNew.Next = &ListNode{Val: itr1.Val}
			itr1 = itr1.Next
		} else {
			itrNew.Next = &ListNode{Val: itr2.Val}
			itr2 = itr2.Next
		}
		itrNew = itrNew.Next
	}

	for itr1 != nil {
		itrNew.Next = &ListNode{Val: itr1.Val}
		itr1 = itr1.Next
		itrNew = itrNew.Next
	}

	for itr2 != nil {
		itrNew.Next = &ListNode{Val: itr2.Val}
		itr2 = itr2.Next
		itrNew = itrNew.Next
	}

	return newHead
}
