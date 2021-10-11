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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	currentNode := res
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			newNode := &ListNode{Val: l1.Val}
			currentNode.Next = newNode
			l1 = l1.Next
			currentNode = newNode
		}else {
			newNode := &ListNode{Val: l2.Val}
			currentNode.Next = newNode
			l2 = l2.Next
			currentNode = newNode
		}
	}
	for l1 != nil{
		newNode := &ListNode{Val: l1.Val}
		currentNode.Next = newNode
		l1 = l1.Next
		currentNode = newNode
	}
	for l2 != nil{
		newNode := &ListNode{Val: l2.Val}
		currentNode.Next = newNode
		l2 = l2.Next
		currentNode = newNode
	}
	return res.Next
}
