package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := NewListNode(-1)
	cur := head
	c := 0
	for l1 != nil || l2 != nil || c == 1 {
		if l1 != nil {
			c += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			c += l2.Val
			l2 = l2.Next
		}
		cur.Next = NewListNode(c % 10)
		cur = cur.Next
		c /= 10
	}

	return head.Next
}
