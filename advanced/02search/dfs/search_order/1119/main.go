// 不重复不漏搜到所有方案
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 21

var (
	n     int
	start string
	words [N]string

	g [N][N]int // g[i][j]=k 表示words[i]的最后k个字符和words[j]的前k个字符是相同的,k是非贪婪的,且小于两者长度的最短值

	cnt [N]int // 表示单词被使用了几次

	ans int // 最终的答案
)

// subStr 返回s[l:r]字串,左闭右开
func subStr(s string, l, r int) string {
	return s[l:r]
}

func dfs(dragon string, last int) {
	ans = max(len(dragon), ans)
	cnt[last]++
	for i := 0; i < n; i++ {
		/*上一个拼接过来的单词和当前遍历到的单词有重合部分*/
		/*每个单词不能使用超过两次*/
		if g[last][i] > 0 && cnt[i] < 2 {
			// 重合部分长度是g[last][i], 那么最后一个重合的下标就是g[last][i]-1,因此不重合的剩余部分的起始下标就是g[last][i]
			dfs(dragon+subStr(words[i], g[last][i], len(words[i])), i)
		}
	}
	cnt[last]--
}

func max(i int, a int) int {
	if i > a {
		return i
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &words[i])
	}
	fmt.Fscanf(reader, "%s\n", &start)

	// 处理得到g数组
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			li := len(words[i])
			lj := len(words[j])
			for k := min(li, lj) - 1; k >= 1; k-- {
				if subStr(words[i], li-k, li) == subStr(words[j], 0, k) {
					g[i][j] = k
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if string(words[i][0]) == start {
			dfs(words[i], i)
		}
	}
	fmt.Println(ans)
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
