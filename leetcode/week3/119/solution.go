package leetcode

func getRow(n int) []int {
	g := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		g[i] = make([]int, i+1)
		g[i][0], g[i][i] = 1, 1
		if i >= 1 {
			for j := 1; j < i; j++ {
				g[i][j] = g[i-1][j-1] + g[i-1][j]
			}
		}
	}
	return g[n]
}
