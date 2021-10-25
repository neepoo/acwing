package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(l, r int, nums []int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = build(l, mid-1, nums)
	root.Right = build(mid+1, r, nums)
	return root
}
func sortedArrayToBST(nums []int) *TreeNode {
	return build(0, len(nums)-1, nums)
}
