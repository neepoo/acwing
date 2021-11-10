package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 25

var (
	g  [N]string
	st [N][N]bool

	sx, sy int

	dx   = []int{-1, 0, 1, 0}
	dy   = []int{0, 1, 0, -1}
	h, w int

	q = [N * N]pair{}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 1000000)
	fmt.Fscanf(reader, "%d %d\n", &w, &h)
	for h != 0 && w != 0 {
		for i := 0; i < h; i++ {
			fmt.Fscanf(reader, "%s\n", &g[i])
			for j := 0; j < len(g[i]); j++ {
				if g[i][j] == '@' {
					sx = i
					sy = j
					break
				}
			}
		}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				st[i][j] = false
			}
		}
		//fmt.Println(dfs(sx, sy))
		fmt.Println(floodFill(sx, sy))
		fmt.Fscanf(reader, "%d %d\n", &w, &h)
	}

}

func dfs(x, y int) int {
	var res int
	st[x][y] = true
	if g[x][y] == '#' {
		return res
	}
	if g[x][y] == '.' || g[x][y] == '@' {
		res++
		for i := 0; i < 4; i++ {
			a := x + dx[i]
			b := y + dy[i]
			if a < 0 || a >= h || b < 0 || b >= w {
				continue
			}
			if st[a][b] {
				continue
			}
			res += dfs(a, b)
		}
	}
	return res
}

type pair struct {
	x, y int
}

func floodFill(x, y int) int {
	var res int
	hh, tt := 0, -1
	tt++
	q[tt] = pair{x, y}
	for hh <= tt{
		head := q[hh]
		hh++

		st[head.x][head.y] = true

		for i := 0; i < 4; i++ {
			a := head.x + dx[i]
			b := head.y + dy[i]

			if a < 0 || a >= h || b < 0 || b >= w{
				continue
			}

			if st[a][b]{
				continue
			}
			if g[a][b] == '#'{
				continue
			}
			res++
			tt++
			st[a][b] = true
			q[tt] = pair{a,b}
		}
	}
	return res + 1
}
