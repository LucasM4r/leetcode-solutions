package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pointer := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pointer.Next = list1
			list1 = list1.Next
		} else {
			pointer.Next = list2
			list2 = list2.Next
		}
		pointer = pointer.Next
	}

	if list1 != nil {
		pointer.Next = list1
	} else {
		pointer.Next = list2
	}

	return dummy.Next
}
