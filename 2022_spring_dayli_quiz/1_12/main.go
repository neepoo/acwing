package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

const N = 50010

var (
	n, k int
	l    = list.New()        // 长度最长为k+1的滑动窗口,里面记录了牛的ID
	s    = make(map[int]int) // 长度为k+1的窗口内每个牛种类出现的次数
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &k)
	var (
		x   int
		res = -1
	)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		if i <= k {
			l.PushBack(x)
			s[x]++
		} else {
			// 窗口满了, 窗口左移
			ele := l.Remove(l.Front())
			// 更新计数
			s[ele.(int)]--
			// 假如窗口
			l.PushBack(x)
			s[x]++
		}
		if s[x] >= 2 {
			// 更新答案
			res = max(res, x)
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
