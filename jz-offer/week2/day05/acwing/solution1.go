package acwing

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Val = -1
	dummy.Next = head
	p := dummy  // p指向上一段的最末
	for p.Next != nil{
		q := p.Next
		//              下一段开头的起始元素
		for q != nil && p.Next.Val == q.Val{
			q = q.Next
		}
		// 循环出来, q指向下一段开头的第一个原始

		// 判断该断的长度
		if p.Next.Next == q{
			// 长度为一, 直接前进一步
			p = p.Next
		}else {
			// 把重复元素这一段全部干掉
			p.Next = q
		}
	}
	return dummy.Next
}
