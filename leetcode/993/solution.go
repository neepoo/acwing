package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	parent, level int
	td            *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	m := make(map[int]node)
	var q [110]node
	hh, tt := 0, -1
	tt++
	q[tt] = node{0, 1, root}
	for hh <= tt {
		head := q[hh]
		hh++
		m[head.td.Val] = node{head.parent, head.level, head.td}
		if head.td.Left != nil {
			tt++
			q[tt] = node{head.td.Val, head.level + 1, head.td.Left}
		}
		if head.td.Right != nil {
			tt++
			q[tt] = node{head.td.Val, head.level + 1, head.td.Right}
		}
	}
	return (m[x].parent != m[y].parent) && (m[x].level == m[y].level)
}
