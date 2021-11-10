package day01

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
	q  [10010]pair
	st [110][110]bool
)

func movingCount(k int, m int, n int) int {
	var res int
	if n == 0 {
		return 0
	}
	floodFill(0, 0, m, n, k, &res)

	return res
}

type pair struct {
	x, y int
}

func floodFill(x, y, m, n, k int, res *int) {
	for i := 0; i < 110; i++ {
		for j := 0; j < 110; j++ {
			st[i][j] = false
		}
	}
	hh, tt := 0, -1
	tt++
	st[x][y] = true
	*res++
	q[tt] = pair{x, y}
	for hh <= tt {
		head := q[hh]
		hh++

		for i := 0; i < 4; i++ {
			a := head.x + dx[i]
			b := head.y + dy[i]
			if a < 0 || a >= m || b < 0 || b >= n {
				continue
			}
			if getSum(a, b) > k {
				continue
			}
			if !st[a][b] {
				*res++
				st[a][b] = true
				tt++
				q[tt] = pair{a, b}
			}
		}
	}

}

func getSum(a, b int) int {
	var ans int
	for a > 0 {
		ans += a % 10
		a /= 10
	}
	for b > 0 {
		ans += b % 10
		b /= 10
	}
	return ans
}
