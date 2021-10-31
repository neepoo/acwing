package solution

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	st := -1
	res := -1
	for i := 0; i < len(s); i++ {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		// 找到非空格开头
		st = i
		for i < len(s) && s[i] != ' ' {
			i++
		}
		res = i - st
		if i >= len(s)-1 {
			break
		}
	}
	return res
}

func lengthOfLastWordNew(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		for s[i] == ' ' {
			i--
		}
		j := i
		for j >= 0 && s[j] != ' ' {
			j--
		}
		return i - j
	}
	return 0
}
