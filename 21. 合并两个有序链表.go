package hot100

// 设置虚拟头节点，尾插法
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	newHead := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newHead.Next = list1
			list1 = list1.Next
		} else {
			newHead.Next = list2
			list2 = list2.Next
		}
		newHead = newHead.Next
	}
	if list1 == nil {
		newHead.Next = list2
	}
	if list2 == nil {
		newHead.Next = list1
	}
	return dummy.Next
}