package main

import "fmt"

const N = 310

type pair struct {
	x, y int
}

var (
	ans int
	st  [N][N]bool
	dx  = []int{-1, 0, 1, 0}
	dy  = []int{0, 1, 0, -1}
	q   [90010]pair
)

func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	for i:=0; i < row; i++ {
		for j:=0; j < col; j++ {
			if !st[i][j] && grid[i][j] == '1'{
				ans++
				floodFill(i, j, grid, row, col)
			}
		}
	}
	clear()
	return ans
}

func clear() {
	// 垃圾leetcode评测机
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			st[i][j] = false
		}
	}
}

func floodFill(i, j int, g [][]byte, row, col int) {
	hh, tt := 0, -1
	tt++
	q[tt] = pair{i, j}
	st[i][j] = true
	for hh <= tt {
		head := q[hh]
		hh++
		x, y := head.x, head.y
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a < 0 || a >= row || b < 0 || b >= col {
				continue
			}
			if st[a][b] {
				continue
			}
			if g[a][b] == '0' {
				continue
			}
			tt++
			st[a][b] = true
			q[tt] = pair{a, b}
		}
	}
}

func main() {
	//b := [][]byte{
	//	{'1', '1', '0', '0', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '1', '0', '0'},
	//	{'0', '0', '0', '1', '1'},
	//}
	//numIslands(b)
	//fmt.Println(ans)

	c := [][]byte{
		{'1'},

	}
	numIslands(c)
	fmt.Println(ans)

	//ans = 0

	//a := [][]byte{
	//	{'1', '1', '1', '1', '0'},
	//	{'1', '1', '0', '1', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '0', '0', '0'},
	//}
	//numIslands(a)
	//fmt.Println(ans)
}
