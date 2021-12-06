package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nthUglyNumber(n int) int {
	q := make([]int, 0, n)
	q[0] = 1
	i, j, k := 0, 0, 0
	for n > 1 {
		t := min(min(q[i]*2, q[j]*3), q[k]*5)
		q = append(q, t)
		if q[i]*2 == t {
			i++
		}
		if q[j]*3 == t {
			j++
		}
		if q[k]*5 == t {
			k++
		}
		n--
	}
	return q[n-1]
}

func main() {

}
