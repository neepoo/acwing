package leetcode

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	i, j := 0, 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s){
		return 0
	}
	// 过滤正负号
	negativeFlag := false
	if s[i] == '+' {
		negativeFlag = false
		i++
	} else if s[i] == '-' {
		negativeFlag = true
		i++
	}
	if i >= len(s){
		//s = "+"
		return 0
	}
	if !isNum(rune(s[i])) {
		return 0
	}
	j = i + 1
	for j < len(s) && isNum(rune(s[j])) {
		j++
	}
	if negativeFlag {
		return convert("-" + s[i:j])
	}
	return convert(s[i:j])
}

func isNum(a rune) bool {
	if a >= '0' && a <= '9' {
		return true
	}
	return false
}

func convert(s string) int {
	res := 0
	negativeFlag := false
	overflow := false
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			negativeFlag = true
		}
		s = s[1:]
	}
	for _, c := range []uint8(s) {
		if !isNum(rune(c)) {
			return 0
		}
		v := c - '0'
		if res > ((1<<31)-1 - int(v)) / 10{
			overflow = true
			break
		}
		res = res*10 + int(v)
	}
	if overflow{
		if negativeFlag{
			return -2147483648
		}else {
			return (1 << 31) - 1
		}
	}
	if negativeFlag {
		if -1*res < -2147483648 {
			return -2147483648
		}
		return -1 * res
	}
	if res > (1<<31)-1 {
		return (1 << 31) - 1
	}
	return res
}
