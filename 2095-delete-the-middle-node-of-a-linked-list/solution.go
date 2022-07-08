package delete_the_middle_node_of_a_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	fast, delayed := head.Next, head
	moveDelayed := false

	if fast == nil {
		return nil
	}

	for fast.Next != nil {
		fast = fast.Next
		if moveDelayed {
			delayed = delayed.Next
		}
		moveDelayed = !moveDelayed
	}

	delayed.Next = delayed.Next.Next
	return head
}
