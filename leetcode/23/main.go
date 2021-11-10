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

func mergeList(p,q *ListNode) *ListNode{
	res := new(ListNode)
	cur := res
	for p != nil && q != nil{
		nd := new(ListNode)
		if p.Val < q.Val{
			nd.Val = p.Val
			p = p.Next
		}else {
			nd.Val = q.Val
			q = q.Next
		}
		cur.Next = nd
		cur = nd
	}
	for p != nil{
		nd := new(ListNode)
		nd.Val = p.Val
		p = p.Next
		cur.Next = nd
		cur = nd
	}
	for q != nil{
		nd := new(ListNode)
		nd.Val = q.Val
		q = q.Next
		cur.Next = nd
		cur = nd
	}
	return res.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeList(lists[0], lists[1])
	default:
		return mergeList(lists[0], mergeKLists(lists[1:]))
	}
}
