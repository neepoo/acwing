package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(a, b *TreeNode, res *[]*TreeNode) {
	if a == nil {
		return
	}
	if a.Val == b.Val {
		*res = append(*res, a)
	}
	dfs(a.Left, b, res)
	dfs(a.Right, b, res)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	var a []*TreeNode
	dfs(A, B, &a)
	for _, nd := range a{
		if equal(nd, B){
			return true
		}
	}
	return false
}

func equal(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	return a.Val == b.Val && equal(a.Left, b.Left) && equal(a.Right, b.Right)
}
func main() {
	A := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	B := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	isSubStructure(A, B)
}
