package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 比较当前两颗子树是否对称
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	// 走到这里说明两颗子树都不同时为空
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
}
