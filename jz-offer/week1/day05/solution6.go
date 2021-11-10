package day05

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Father *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	return _buildTree(preorder, inorder, len(preorder))
}

// l表示序列的长度
func _buildTree(preorder []int, inorder []int, l int) *TreeNode {
	if l == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	// idx左边为左子树,右边为右子树
	var idx int
	for inorder[idx] != root.Val {
		idx++
	}
	root.Left = _buildTree(preorder[1:], inorder, idx)
	root.Right = _buildTree(preorder[idx+1:], inorder[idx+1:], l-idx-1)
	return root
}
