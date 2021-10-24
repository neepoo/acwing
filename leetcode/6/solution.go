package leetcode

func convert(s string, numRows int) string {
	if numRows == 1{
		// 需要特殊判断,否则 x++会出错
		return s
	}
	g := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		g[i] = make([]string, len(s))
		for j := 0; j < len(s); j++ {
			g[i][j] = "/"
		}
	}
	x, y := 0, 0
	down, up := true, false
	for _, c := range s {
		g[x][y] = string(c)
		if x == numRows-1 {
			down = false
			up = true
		}
		if x == 0 {
			down = true
			up = false
		}
		if down {
			x++
		}
		if up {
			x--
			y++
		}
	}
	res := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			if g[i][j] != "/"  {
				res += g[i][j]
			}
		}
	}
	return res
}
