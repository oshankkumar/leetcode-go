package p2_add_two_numbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	curr *ListNode
}

func (l *List) Push(val int) {
	if l.Head == nil {
		l.Head = &ListNode{Val: val}
		l.curr = l.Head
		return
	}

	l.curr.Next = &ListNode{Val: val}
	l.curr = l.curr.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumList List

	var carry int
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		nodeVal := sum % 10
		carry = sum / 10
		sumList.Push(nodeVal)

		l1, l2 = l1.Next, l2.Next
	}

	addRemaining := func(l *ListNode) {
		for l != nil {
			sum := l.Val + carry
			nodeVal := sum % 10
			carry = sum / 10
			sumList.Push(nodeVal)
			l = l.Next
		}
	}

	addRemaining(l1)
	addRemaining(l2)

	if carry != 0 {
		sumList.Push(carry)
	}

	return sumList.Head
}
