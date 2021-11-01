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

func printListReversingly(head *ListNode) []int {
	if head == nil{
		return []int{}
	}
	if head.Next == nil{
		return []int{head.Val}
	}
	return append(printListReversingly(head.Next), head.Val)
}
