package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var ans []int
	g := make(map[*TreeNode][]*TreeNode)
	buildGraph(root, g)
	dfs(target, nil, k, &ans, g)
	return ans
}

func buildGraph(root *TreeNode, m map[*TreeNode][]*TreeNode) {
	if root.Left != nil {
		m[root] = append(m[root], root.Left)
		m[root.Left] = append(m[root.Left], root)
		buildGraph(root.Left, m)
	}

	if root.Right != nil {
		m[root] = append(m[root], root.Right)
		m[root.Right] = append(m[root.Right], root)
		buildGraph(root.Right, m)
	}
}

func dfs(root *TreeNode, father *TreeNode,k int, ans *[]int, m map[*TreeNode][]*TreeNode){
	if k == 0{
		*ans = append(*ans, root.Val)
	}else {
		for _, node := range m[root]{
			if node != father{
				// 避免走回头路
				dfs(node, root, k-1, ans, m)
			}
		}
	}
}
