package leetcode



func romanToInt(s string) int {
	var m = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		switch rune(s[i]) {
		case 'I':
			if i+1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X') {
				res += m[string(rune(s[i+1]))] - 1
				i++
			} else {
				res += 1
			}
		case 'V':
			res += 5
		case 'X':
			if i+1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C') {
				res += m[string(rune(s[i+1]))] - 10
				i++
			} else {
				res += 10
			}
		case 'L':
			res += 50
		case 'C':
			if i+1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M') {
				res += m[string(rune(s[i+1]))] - 100
				i++
			} else {
				res += 100
			}
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}
	}
	return res
}
