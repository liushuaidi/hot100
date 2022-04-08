package hot100

// 模拟题
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	res := cur
	var sum, v, t int = 0, 0, 0
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + t
		v = sum % 10
		t = sum / 10
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum = l1.Val + t
		v = sum % 10
		t = sum / 10
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum = l2.Val + t
		v = sum % 10
		t = sum / 10
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		l2 = l2.Next
	}
	if t > 0 {
		cur.Next = &ListNode{Val: t}
	}
	return res.Next
}

// 优化版
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var cur *ListNode
	res := cur
	var sum, t int = 0, 0
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum = a + b + t
		a = sum % 10
		t = sum / 10
		if cur == nil {
			cur = &ListNode{Val: a}
			res = cur
		} else {
			cur.Next = &ListNode{Val: a}
			cur = cur.Next
		}
	}

	if t > 0 {
		cur.Next = &ListNode{Val: t}
	}
	return res
}
