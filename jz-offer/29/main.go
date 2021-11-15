package main

import "fmt"

func spiralOrder(nums [][]int) []int {
	var res []int
	r := len(nums)
	if r == 0 {
		return res
	}
	c := len(nums[0])
	if c == 0 {
		return res
	}
	st := [110][110]bool{}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	x, y := 0, 0
	st[x][y] = true
	var turn int
	res = append(res, nums[x][y])
	for len(res) < r*c {
		a := x + dx[turn%4]
		b := y + dy[turn%4]
		if a < 0 || a >= r || b < 0 || b >= c || st[a][b] {
			turn++
			a = x
			b = y
			continue
		}
		st[a][b] = true
		res = append(res, nums[a][b])
		x = a
		y = b

	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
