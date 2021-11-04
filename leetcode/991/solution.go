package leetcode

func brokenCalc(s int, t int) int {
	var res int
	for t > s {
		if t%2 == 1 {
			t++
		} else {
			t /= 2
		}
		res++
	}
	return res + s - t
}
