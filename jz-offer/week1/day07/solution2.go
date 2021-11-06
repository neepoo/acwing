package main

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func hasPath(matrix [][]byte, str string) bool {
	r := len(matrix)
	if r == 0 {
		return false
	}
	c := len(matrix[0])
	if c == 0 {
		return false
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if dfs(matrix, str, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(matrix [][]byte, str string, u, x, y int) bool {
	if matrix[x][y] != str[u] {
		return false
	}
	if u == len(str)-1 {
		return true
	}
	t := matrix[x][y]
	matrix[x][y] = '*'
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a >= 0 && a < len(matrix) && b >= 0 && b < len(matrix[0]) {
			if dfs(matrix, str, u+1, a, b) {
				return true
			}
		}
	}
	matrix[x][y] = t
	return false
}

func main() {

}
