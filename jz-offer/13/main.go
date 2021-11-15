package main

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

type pair struct {
	x, y int
}

func movingCount(m int, n int, k int) int {
	var st [101][101]bool
	var res int
	var q [100]pair
	hh, tt := 0, 1
	q[0] = pair{0, 0}
	res++
	st[0][0] = true
	for hh != tt {
		head := q[hh]
		hh++
		if hh == 100 {
			hh = 0
		}
		x, y := head.x, head.y
		for i := 0; i < 4; i++ {
			a := x + dx[i]
			b := y + dy[i]
			if a < 0 || a >= m || b < 0 || b >= n {
				continue
			}
			if st[a][b] {
				continue
			}
			if get(a)+get(b) > k {
				continue
			}
			q[tt] = pair{a, b}
			tt++
			if tt == 100 {
				tt = 0
			}
			st[a][b] = true
			res++

		}
	}
	return res
}

func get(x int) int {
	var res int
	for x > 0 {
		res += x % 10
		x /= 10
		// 之前写成这样了...
		//res /= 10
	}
	return res
}

func main() {
	println(movingCount(2, 3, 1))
	println(movingCount(3, 1, 0))
}
