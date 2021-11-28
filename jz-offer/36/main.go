package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	pre, head *TreeNode
)
func convert(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}

	dfs(root)
	return head
}
func dfs(cur *TreeNode) {
	// left, mid, right
	if cur == nil{
		return
	}
	dfs(cur.Left)
	if pre == nil{
		head = cur
	}else {
		pre.Right = cur
		cur.Left = pre
	}
	pre = cur
	dfs(cur.Right)
}

func main() {

}
