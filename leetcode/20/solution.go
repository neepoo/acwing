package leetcode

func isValid(s string) bool {
	var stk = make([]int32, 10010)
	tt := 0
	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			tt++
			stk[tt] = ch
		case ')':
			left := stk[tt]
			tt--
			if left != '(' {
				return false
			}
		case '}':
			left := stk[tt]
			tt--
			if left != '{' {
				return false
			}
		case ']':
			left := stk[tt]
			tt--
			if left != '[' {
				return false
			}
		}
	}

	return tt == 0 // 栈为空才合法
}
