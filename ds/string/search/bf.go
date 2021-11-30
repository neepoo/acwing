package search

func bs1(s, p string) int {
	n, m := len(s), len(p)
	for i := 0; i <= n-m; i++ {
		j := 0
		for ; j < m; j++ {
			if p[j] != s[i+j] {
				break
			}
		}
		if j == m {
			return i
		}
	}
	return -1
}

// 有启发意义
func bs2(s, p string) int {
	n, m := len(s), len(p)
	i, j := 0, 0 // i:已经匹配字符序列的末端, j当前p串字符的索引
	for ; i < n && j < m; i++ {
		if s[i] == p[j] {
			j++
		} else {
			/*
					如果 i 和 j 指向的字符不匹配了,那么需要回
					退这两个指针的值:将 j 重新指向模式的开头,将 i 指向本次匹配的开始位置的下一个字符。

				kmp的改进,就是能够使得i不回退,只用设置j的值
			*/

			i -= j // 回退到上次匹配的开头,然后重新进入循环,执行i++时,指向上次匹配的下一个字符
			j = 0
		}
	}
	if j == m {
		return i - m
	}
	return -1
}

