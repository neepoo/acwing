package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(h *Node) *Node {

	h1 := h
	h2 := h
	h3 := h
	for h1 != nil {
		nd := new(Node)
		nd.Val = h1.Val
		next := h1.Next
		h1.Next = nd
		nd.Next = next
		h1 = next
	}

	for ;h2 != nil;h2=h2.Next.Next  {
		if h2.Random != nil {
			h2.Next.Random = h2.Random.Next
		}
	}
	dummy := new(Node)
	cur := dummy
	for ;h3 != nil;h3 = h3.Next{
		cur.Next = h3.Next
		cur = cur.Next
		h3.Next = h3.Next.Next
	}
	return dummy.Next
}
func main() {
	n1 := &Node{Val: 7}
	n2 := &Node{Val: 13}
	n3 := &Node{Val: 11}
	n4 := &Node{Val: 10}
	n5 := &Node{Val: 1}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n2.Random = n1
	n3.Random = n5
	n4.Random = n3
	n5.Random = n1

	res := copyRandomList(n1)
	for res != nil {
		fmt.Printf("%#v\n", res)
		res = res.Next
	}
}
