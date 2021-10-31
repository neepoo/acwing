package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur, ne := head, head
	res := cur
	for ne != nil {
		for ne != nil && cur.Val == ne.Val {
			ne = ne.Next
		}
		cur.Next = ne
		cur = ne
	}
	return res
}
