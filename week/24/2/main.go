package main

import (
	"fmt"
)

var (
	x1, y1 int
	x2, y2 int
	dx     = []int{-2, -1, 1, 2, 2, 1, -1, -2}
	dy     = []int{1, 2, 2, 1, -1, -2, -2, -1}
)

func check1(x, y int) bool {
	var res bool
	if x == x1 || y == y1 {
		res = true
	}
	return res
}

// 马是互相攻击
func check2(h1x, h1y, h2x, h2y int) bool {
	var res bool
	for i := 0; i < 8; i++ {
		a := h1x + dx[i]
		b := h1y + dy[i]
		if a < 1 || a > 8 || b < 1 || b > 8 {
			continue
		}
		if a == h2x && b == h2y {
			res = true
		}
	}
	return res
}

func main() {
	var s1, s2 string
	fmt.Scanf("%s\n%s", &s1, &s2)
	y1 = int(s1[0]-'a') + 1
	y2 = int(s2[0]-'a') + 1

	x1 = int(s1[1] - '1')
	x2 = int(s2[1] - '1')
	x2 = 8 - x2
	x1 = 8 - x1
	var res int
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if (x1 == i && y1 == j) || (x2 == i && y2 == j) {
				continue
			}
			if !check1(i, j) && !check2(i, j, x2, y2) && !check2(i, j, x1, y1) {
				res++
			}
		}
	}
	fmt.Println(res)
}
