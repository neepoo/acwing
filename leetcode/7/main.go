package leetcode

import "math"

func reverse(x int) int {
	var res int32 = 0
	negativeFlag := false
	if x < 0 {
		negativeFlag = true
		x = -x
	}
	for x > 0 {
		m := int32(x % 10)
		// 判断溢出
		if res > (math.MaxInt32-m)/10 {
			return 0
		}
		res = res*10 + m
		x /= 10
	}
	if negativeFlag {
		return -1 * int(res)
	}
	return int(res)
}
