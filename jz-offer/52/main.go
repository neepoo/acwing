package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	l1, l2 := 0, 0
//	h1, h2 := headA, headB
//	for h1 != nil {
//		l1++
//		h1 = h1.Next
//	}
//	for h2 != nil {
//		l2++
//		h2 = h2.Next
//	}
//	// 保证 h1是长度较长的那个链表
//	if l1 < l2 {
//		l1, l2 = l2, l1
//		headA, headB = headB, headA
//	}
//	for l1 > l2 {
//		headA = headA.Next
//		l1--
//	}
//	k := 0
//	for k < l2 {
//		if headA == headB {
//			return headA
//		}
//		k++
//		headA = headA.Next
//		headB = headB.Next
//	}
//	return nil
//
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	h1, h2 := headA, headB
	for h1 != h2 {
		if h1 == nil {
			h1 = headB
		} else {
			h1 = h1.Next
		}
		if h2 == nil {
			h2 = headA
		} else {
			h2 = h2.Next
		}
	}
	return h1
}

func main() {

}
