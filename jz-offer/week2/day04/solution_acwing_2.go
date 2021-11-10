package day04


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func deleteNode(node *ListNode)  {
	// 复制下一个节点到当前,删除下一个节点
	ne := node.Next
	node.Val = ne.Val
	node.Next = ne.Next
	ne = nil
}