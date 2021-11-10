package day07

func entryNodeOfLoop(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	p := head
	for p != nil{
		m[p]++
		if m[p] >1 {
			return p
		}
		p = p.Next
	}
	return nil
}
