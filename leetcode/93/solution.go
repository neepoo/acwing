package solution

import "strconv"

func restoreIpAddresses(s string) []string {
	ans := new([]string)
	n := len(s)
	var dfs func(int, int, string, *[]string)
	// u表示枚举到第几位, t表示枚举的第几个数(从0开始), cur表示当前的ip表示, ans表示最终的答案
	dfs = func(u, t int, cur string, ans *[]string) {
		if u == n && t == 4 {
			*ans = append(*ans, cur[:len(cur)-1])
			return
		}
		// 减枝
		if t == 4 {
			return
		}
		k := 0
		for i := u; i < n; i++ {
			// 前导0
			if i > u && s[u] == '0' {
				continue
			}
			k = k*10 + int(s[i]-'0')
			if k <= 255 {
				s2 := cur + strconv.Itoa(k) + "."
				dfs(i+1, t+1, s2, ans)
			}
		}
	}
	dfs(0, 0, "", ans)
	return *ans
}
