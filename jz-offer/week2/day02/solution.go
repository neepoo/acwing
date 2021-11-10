package day02

func maxProductAfterCutting(n int) int {
	var res = 1
	if n <= 3 {
		return n - 1
	}
	if n%3 == 1 {
		res = 4
		n -= 4
	} else if n%3 == 2 {
		res = 2
		n -= 2
	}
	for n > 0 {
		res *= 3
		n -= 3
	}

	return res
}

func dp(n int) int {
	var f [60]int
	f[2], f[3] = 1, 2
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			f[i] = max(f[i], max((i-j)*j, (i-j)*f[j]))
		}
	}
	return f[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
