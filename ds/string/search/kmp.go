package search

// i一直往后,不匹配时,j没必要每次从0开始,只要s[i]往前的k个字符和p串开头的k个字符相同就可以
// j回退到k就可以
// 怎么才能知道p的开头和s[i]之前的字符相同?
func kmp(s, p string) int {
	n, m := len(s), len(p)
	getNext := func() []int {
		next := make([]int, m+1)
		next[0] = -1
		j, k := 1, -1
		for j < m {
			if k == -1 || p[j] == p[k] {
				j++
				k++
				next[j] = k
			} else {
				k = next[k]
			}
		}

		return next
	}

	next := getNext()
	i, j := 0, 0
	for ; i < n && j < m; i++ {
		if s[i] != p[j] {
			j = next[j]
		}
		j++
	}
	if j == m {
		return i - m
	} else {
		return -1
	}

}
