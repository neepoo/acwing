// https://www.acwing.com/problem/content/1989/
// 差分 + 离散化
// 左端点映射成一个区间
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b := make(map[int]int)
	var n, x, y int
	var op string
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &y, &op)
		if op == "R" {
			b[x]++
			b[x+y]--
			x += y
		} else {
			b[x]--
			b[x-y]++
			x -= y
		}
	}
	keys := make([]int, len(b))
	i := 0
	for k := range b {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	var res, sum, last int
	for _, k := range keys {
		key, val := k, b[k]
		if sum >= 2 {
			res += key - last
		}
		sum += val
		last = key
	}

	fmt.Println(res)
}
