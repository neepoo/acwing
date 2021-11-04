package leetcode

const (
	R      = 'R'
	Bishop = 'B' //白色的象
	p      = 'p'
	dot    = '.'
)

func numRookCaptures(board [][]byte) int {
	var (
		sx, sy int

		dx = []int{-1, 0, 1, 0}
		dy = []int{0, 1, 0, -1}
	)
	rl, cl := len(board), len(board[0])
	for i := 0; i < rl; i++ {
		for j := 0; j < cl; j++ {
			if board[i][j] == R {
				sx, sy = i, j
			}
		}
	}

	var res int
	for i := 0; i < 4; i++ {
		r, l := sx, sy
		// 棋盘边缘不合法
		// 遇到Bishop不合法
		// 吃了pawn res++, 也要停下来
		for {
			r += dx[i]
			l += dy[i]
			if r < 0 || r >= rl || l < 0 || l >= cl {
				break
			}
			if board[r][l] == dot{
				continue
			}
			if board[r][l] == Bishop {
				break
			}
			if board[r][l] == p {
				res++
				break
			}
		}
	}
	return res
}
