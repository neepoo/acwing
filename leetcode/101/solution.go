package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(q, p *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil || p.Val != q.Val {
		return false
	}
	return dfs(q.Left, p.Right) && dfs(q.Right, p.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return dfs(root.Left, root.Right)
}
