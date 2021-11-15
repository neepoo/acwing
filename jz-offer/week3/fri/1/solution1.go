package main

import "fmt"

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

type pair struct {
	nd    *TreeNode
	depth int
}

func printFromTopToBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		q      [1000]pair
		res    [10000][]int
		hh, tt = 0, 1
	)
	dp := 0
	q[0] = pair{root, 0}
	for hh != tt {
		head := q[hh]
		hh++
		if hh == 1000 {
			hh = 0
		}
		dp = max(dp, head.depth)
		res[head.depth] = append(res[head.depth], head.nd.Val)
		if head.nd.Left != nil {
			q[tt] = pair{head.nd.Left, head.depth + 1}
			tt++
			if tt == 1000 {
				tt = 0
			}
		}
		if head.nd.Right != nil {
			q[tt] = pair{head.nd.Right, head.depth + 1}
			tt++
			if tt == 1000 {
				tt = 0
			}
		}

	}
	return res[:dp+1]
}

func max(dp int, depth int) int {
	if dp > depth {
		return dp
	}
	return depth
}


func main() {
	fmt.Println(printFromTopToBottom(&TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   12,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}
