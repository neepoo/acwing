package day07

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *
 *
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func findKthToTail(head *ListNode, k int) *ListNode {
	p, q := head, head
	var n int
	for q != nil{
		n++
		q = q.Next
	}
	if k > n{
		return nil
	}
	var i = n - k
	for i > 0{
		p = p.Next
		i--
	}
	return p
}
