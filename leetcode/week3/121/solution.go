package leetcode

func maxProfit(prices []int) int {
	var mp = 0x3f3f3f3f
	var ans = 0
	for _, v := range prices {
		ans = max(ans, v-mp)
		mp = min(mp, v)
	}
	return ans
}

func max(ans int, i int) int {
	if ans > i {
		return ans
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
