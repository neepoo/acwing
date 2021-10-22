package solution

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	st := -1
	res := -1
	for i := 0; i < len(s); i++ {
		for i < len(s) && s[i] == ' '{
			i++
		}
		// 找到非空格开头
		st = i
		for i < len(s) && s[i]!=' '{
			i++
		}
		res = i - st
		if i >= len(s) - 1{
			break
		}
	}
	return res
}
