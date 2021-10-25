package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(depth int, depth2 int) int {
	if depth > depth2 {
		return depth
	}
	return depth2
}