package leetcode

func generate(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		if i >= 1 {
			// 每行数组的下标范围时从0 - i
			for j := 1; j < i; j++ {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}

	}
	return res
}
