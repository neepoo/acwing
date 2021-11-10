package leetcode



type ListNode struct {
	Val int
	Next *ListNode
}
func deleteNode(head *ListNode, val int) *ListNode {
	cur := head
	for head.Next != nil && head.Next.Val != val {
			head = head.Next
	}
	// head.Next就是要删除的节点
	head.Next = head.Next.Next
	return cur
}