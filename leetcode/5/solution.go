package leetcode

// 思路: 枚举中间点,同时往两边走
func longestPalindrome(s string) string {
	L := len(s)
	res := ""
	for i := 0; i < L; i++ {
		// 枚举回文串长度时奇数的情况
		l, r := i-1, i+1
		for l >= 0 && r <= L-1 && s[l] == s[r] {
			l--
			r++
		}
		if r-1-(l+1)+1 > len(res) {
			res = s[l+1 : r]
		}
		// 枚举回文串长度是偶数的情况
		l, r = i, i+1
		for l >= 0 && r <= L-1 && s[l] == s[r] {
			l--
			r++
		}
		if r-1-(l+1)+1 > len(res) {
			res = s[l+1 : r]
		}
	}
	return res
}
