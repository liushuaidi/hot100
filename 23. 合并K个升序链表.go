package hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1、顺序合并
func mergeKLists01(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists(ans, lists[i])
	}
	return ans
}

// 2、分治合并
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := left + (right-left)>>1
	return mergeTwoLists(merge(lists, left, mid), merge(lists, mid, right))
}

// ****************************************

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
