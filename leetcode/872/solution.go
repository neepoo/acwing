package leetcode

import (
	"reflect"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(r *TreeNode) bool {
	return r.Left == nil && r.Right == nil
}

func dfs(r *TreeNode, ns *[]int) {
	if r == nil{
		return
	}
	if isLeaf(r){
		*ns = append(*ns, r.Val)
	}
	if r.Left != nil{
		dfs(r.Left, ns)
	}
	if r.Right != nil{
		dfs(r.Right, ns)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var nums1 []int
	var nums2 []int
	dfs(root1, &nums1)
	dfs(root2, &nums2)
	return reflect.DeepEqual(nums1, nums2)
}
