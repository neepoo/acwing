package leetcode

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	const (
		N = 25
		M = 35
	)
	var f [N][M]bool
	s = " " + s
	p = " " + p
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j+1 <= m && p[j+1] == '*' {
				continue
			}
			if i > 0 && p[j] != '*' {
				f[i][j] = f[i-1][j-1] && (s[i] == p[j] || p[j] == '.')
			} else if p[j] == '*' {
				f[i][j] = f[i][j-2] || (i > 0 && (f[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.')))
			}
		}
	}
	return f[n][m]
}
