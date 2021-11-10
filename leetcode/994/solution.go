package leetcode

func orangesRotting(grid [][]int) int {
	const (
		bad  = 2
		good = 1
		INF  = 0x3f3f3f3f
	)
	type pair struct {
		x, y, d int
	}
	var (
		q  [110]pair
		dx = []int{-1, 0, 1, 0}
		dy = []int{0, 1, 0, -1}
	)
	hh, tt := 0, -1
	l, r := len(grid), len(grid[0])
	dist := make([][]int, l)
	isGood := make([][]bool, l)
	st := make([][]bool, l)
	goodCnt := 0
	for i := 0; i < l; i++ {
		dist[i] = make([]int, r)
		isGood[i] = make([]bool, r)
		st[i] = make([]bool, r)
		for j := 0; j < r; j++ {
			if grid[i][j] == good {
				goodCnt++
				isGood[i][j] = true
				dist[i][j] = INF
			} else if grid[i][j] == bad {
				tt++
				q[tt] = pair{i, j, 0}
				st[i][j] = true
			}
		}
	}
	if goodCnt == 0{
		return 0
	}
	for hh <= tt {
		head := q[hh]
		hh++
		for i := 0; i < 4; i++ {
			a := head.x + dx[i]
			b := head.y + dy[i]
			if a < 0 || a >= l || b < 0 || b >= r {
				continue
			}
			if st[a][b] {
				continue
			}
			if isGood[a][b] {
				// 更新距离
				dist[a][b] = min(dist[a][b], dist[head.x][head.y]+1)
				tt++
				q[tt] = pair{a, b, head.d + 1}
				st[a][b] = true
			}
		}
	}
	var res = -INF
	for i := 0; i < l; i++ {
		for j := 0; j < r; j++ {
			if isGood[i][j]{
				res = max(res, dist[i][j])
			}
		}
	}
	if res == INF{
		return -1
	}
	return res
}

func min(i int, i2 int) int {
	if i < i2{
		return i
	}
	return i2
}

func max(i int, i2 int) int {
	if i > i2{
		return i
	}
	return i2
}