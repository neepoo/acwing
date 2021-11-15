package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	nd    *TreeNode
	depth int
	left  bool
}
func printFromTopToBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var (
		q      [100]pair
		res    [1000][]int
		hh, tt = 0, 1
	)
	dp := 0
	q[0] = pair{root, 0, true}
	for hh != tt {
		head := q[hh]
		hh++
		if hh == 100 {
			hh = 0
		}
		dp = max(dp, head.depth)
		if head.left {
			res[head.depth] = append(res[head.depth], head.nd.Val)
			if head.nd.Right != nil {
				q[tt] = pair{head.nd.Right, head.depth + 1, !head.left}
				tt++
				if tt == 100 {
					tt = 0
				}
			}
			if head.nd.Left != nil {
				q[tt] = pair{head.nd.Left, head.depth + 1, !head.left}
				tt++
				if tt == 100 {
					tt = 0
				}
			}
		}
		if !head.left {
			res[head.depth] = append(res[head.depth], head.nd.Val)
			if head.nd.Left != nil {
				q[tt] = pair{head.nd.Left, head.depth + 1, !head.left}
				tt++
				if tt == 100 {
					tt = 0
				}
			}
			if head.nd.Right != nil {
				q[tt] = pair{head.nd.Right, head.depth + 1, !head.left}
				tt++
				if tt == 100 {
					tt = 0
				}
			}
		}
	}
	for i:=0;i<dp+1;i++{
		if i % 2 == 0{

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
