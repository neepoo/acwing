package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans = true

func isBalanced(root *TreeNode) bool {
	ans = true
	if root == nil {
		return true
	}
	dfs(root)
	return ans
}

// 返回以root为根节点子树的最大高度
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl, hr := dfs(root.Left), dfs(root.Right)
	if abs(hl-hr) > 1 {
		ans = false
	}
	return 1 + max(hl, hr)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -1 * a
}
