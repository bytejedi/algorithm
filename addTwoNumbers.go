package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b, current := l1, l2, &ListNode{}
	result := current
	var carry int
	for a != nil || b != nil {
		current.Next = &ListNode{}
		current = current.Next
		var x, y int
		if a != nil {
			x = a.Val
			a = a.Next
		}
		if b != nil {
			y = b.Val
			b = b.Next
		}

		sum := carry + x +y
		carry = sum / 10
		if carry > 0 {
			current.Val = sum - carry*10
		} else {
			current.Val = sum
		}
	}
	if carry > 0 {
		current.Next = &ListNode{Val:carry}
	}
	return result.Next
}
