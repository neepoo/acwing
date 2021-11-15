package main

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
	dummy := new(ListNode)
	cur := dummy
	for l1 != nil && l2 != nil {
		nd := new(ListNode)
		if l1.Val < l2.Val {
			nd.Val = l1.Val
			l1 = l1.Next
		} else {
			nd.Val = l2.Val
			l2 = l2.Next
		}
		cur.Next = nd
		cur = nd
	}
	for l1 != nil {
		nd := new(ListNode)
		nd.Val = l1.Val
		cur.Next = nd
		cur = nd
		l1 = l1.Next

	}
	for l2 != nil {
		nd := new(ListNode)
		nd.Val = l2.Val
		cur.Next = nd
		cur = nd
		l2 = l2.Next
	}
	return dummy.Next
}

func main() {

}
