package leetcode

func maxCount(m int, n int, ops [][]int) int {
	x, y := m, n
	for _, nums := range ops {
		x = min(x, nums[0])
		y = min(y, nums[1])
	}
	return x * y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
