package leetcode

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[uint8]int)
	j := 0
	for i, c := range []uint8(s){
		m[c]++
		for m[c] > 1{
			m[s[j]]--
			j++
		}
		res = max(res, i - j + 1)
	}
	return res
}
