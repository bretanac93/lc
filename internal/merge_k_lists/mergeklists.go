package merge_k_lists

/*
1. Consider the first list as the result, and store it in a variable.
2. Traverse the list of lists, starting from the second list, and combine it with the list we stored as a result.
The result should get stored in the same variable.
3. When combining the two lists, like l1 and l2, maintain a prev pointer that points to a dummy node.
4. If the value for list l1 is less than or equal to the value for list l2, connect the previous node to l1 and increment l1.
Otherwise, do the same but for list l2.
5. Keep repeating the above step until one list points to a nil value.
6. Connect the non-nil list to the merged one and return.
*/

func mergeLists(l1, l2 *LinkedListNode) *LinkedListNode {
	dummy := &LinkedListNode{data: -1}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			prev.next = l1
			l1 = l1.next
		} else {
			prev.next = l2
			l2 = l2.next
		}
		prev = prev.next
	}
	if l1 != nil {
		prev.next = l1
	} else {
		prev.next = l2
	}

	return dummy.next
}

func mergeKLists(lists []*LinkedListNode) *LinkedListNode {
	res := lists[0]
	if len(lists) == 0 {
		return &LinkedListNode{data: -1}
	}
	for i := 1; i < len(lists); i++ {
		res = mergeLists(res, lists[i])
	}
	return res
}
