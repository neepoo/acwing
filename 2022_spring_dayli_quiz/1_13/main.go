/*
https://www.acwing.com/problem/content/1962/
*/
package main

import "fmt"

const N = 1 << 16

var (
	n     int
	m     uint64
	state uint64
	p     [N]uint64 // p[state] = i 表示state状态是第i个得到的
)

func init() {
	for i := 0; i < N; i++ {
		p[i] = 0x3f3f3f3f
	}
}

func pr(x uint64) {
	for i := 0; i < n; i++ {
		fmt.Println(x >> i & 1)
	}
}
func main() {
	fmt.Scan(&n, &m)
	var x uint64
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		state |= x << i
	}
	p[state] = 0
	var i uint64
	for i = 1; ; i++ {
		state = update(state)
		switch {
		case i == m:
			pr(state)
			return
		case p[state] == 0x3f3f3f3f:
			p[state] = i
		default:
			circleLen := i - p[state]
			// 注意这里要减去已经走了的步数
			r := (m - i) % circleLen
			for r > 0 {
				state = update(state)
				r--
			}
			pr(state)
			return
		}
	}
}

func update(state uint64) uint64 {
	var res uint64 = 0
	for i := 0; i < n; i++ {
		currentBit := (state >> i) & 1
		preBit := (state >> ((i - 1 + n) % n)) & 1
		x := currentBit ^ preBit
		res |= x << i
	}
	return res
}
