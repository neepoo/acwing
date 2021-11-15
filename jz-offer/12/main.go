package main

type pair struct {
	x, y, i int
}

func exist(board [][]byte, word string) bool {
	r, c := len(board), len(board[0])
	var (
		dx = []int{-1, 0, 1, 0}
		dy = []int{0, 1, 0, -1}
		ps []pair
	)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				ps = append(ps, pair{i, j, 0})
			}
		}
	}
	if len(word) == 1 && len(ps) > 0 {
		return true
	}
	q := [400010]pair{}
	for _, s := range ps {
		st := [110][110]bool{}
		hh, tt := 0, -1
		tt++
		q[tt] = s
		st[s.x][s.y] = true
		for hh <= tt {
			head := q[hh]
			hh++
			x, y, l := head.x, head.y, head.i
			st[x][y] = false
			for i := 0; i < 4; i++ {
				a := x + dx[i]
				b := y + dy[i]
				newL := l + 1
				if a < 0 || a >= r || b < 0 || b >= c {
					continue
				}
				if st[a][b]{
					continue
				}

				if board[a][b] == word[newL]{
					tt++
					st[a][b]=true
					q[tt] = pair{a, b, newL}
					if newL == len(word)-1 {
						return true
					}
				}
			}
		}
	}

	return false
}
func main() {

}
