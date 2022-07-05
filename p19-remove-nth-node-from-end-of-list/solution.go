package p19

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr, delayed := head, head

	for i := 1; i <= n; i++ {
		curr = curr.Next
	}

	if curr == nil {
		return head.Next
	}

	for curr.Next != nil {
		curr, delayed = curr.Next, delayed.Next
	}

	delayed.Next = delayed.Next.Next

	return head
}
