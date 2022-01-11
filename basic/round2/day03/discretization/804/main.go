// https://www.acwing.com/problem/content/804/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

type ps []pair

const N = 3e5 + 10

var (
	n, m int
	x, c int
	l, r int
	all  []int // 映射之后的数组
	//
	add   ps
	query ps
	s     [N]int
)

func (p ps) Len() int {
	return len(p)
}

func (p ps) Less(i, j int) bool {
	return p[i].x < p[j].x
}

func (p ps) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	all = make([]int, 0, N)
	add = make(ps, 0, 1e5)
	query = make(ps, 0, 1e5)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x, &c)
		add = append(add, pair{x, c})
		all = append(all, x)
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &l, &r)
		query = append(query, pair{l, r})
		all = append(all, l, r)
	}
	sort.Ints(all)
	all = unique(all)
	for _, addOp := range add {
		idx := find(addOp.x)
		s[idx] += addOp.y
	}
	for i := 1; i < len(s); i++ {
		s[i] += s[i-1]
	}
	for _, q := range query {
		l := find(q.x)
		r := find(q.y)
		fmt.Println(s[r] - s[l-1])
	}
}

func find(x int) int {
	l, r := 0, len(all)-1
	for l < r {
		mid := (l + r) / 2
		if all[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r + 1

}

func unique(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	cnt := 1
	last := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] != last {
			nums[cnt] = nums[i]
			last = nums[i]
			cnt++
		}
	}
	return nums[:cnt]
}
