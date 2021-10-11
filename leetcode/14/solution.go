package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	res := ""
	for _, s := range strs[1:] {
		minLen = min(minLen, len(s))
	}
	var lastIndex int = -1
	for i := 0; i < minLen; i++ {
		equalFlag := true
		lastCh := strs[0][i]
		for _, s := range strs[1:] {
			if s[i] != lastCh {
				equalFlag = false
			} else {
				lastCh = s[i]
			}
		}
		if equalFlag {
			if lastIndex == -1 {
				lastIndex = 0
			}
			if i != lastIndex+1 && i != 0{
				return res
			} else {
				res += string(rune(lastCh))
				lastIndex = i
			}
		}else {
			return res
		}
	}
	return res
}
