package solution

func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	if x == 1{
		return 1
	}
	l, r := 0, x
	var mid int
	for k := 0; k <= 31; k++ {
		mid = (l + r) / 2
		i := mid * mid
		if i > x {
			r = mid
		} else if i < x {
			l = mid
		} else {
			return mid
		}
	}
	return mid
}
