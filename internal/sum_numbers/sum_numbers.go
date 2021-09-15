package sumnumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil || l2 != nil || carry != 0 {
		var v1, v2 int

		if l1 != nil {
			v1 = l1.Val
		}

		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2 + carry
		digit := sum % 10
		carry = sum / 10

		curr.Next = &ListNode{digit, nil}

		curr = curr.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	return dummy.Next
}
