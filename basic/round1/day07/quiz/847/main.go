package main

import "fmt"

const N = 3

var (
	p      [N][N]string
	ix, iy int
	dx     = [4]int{-1, 0, 1, 0}
	dy     = [4]int{0, 1, 0, -1}
)

func main() {
	var s string
	for i := 0; i < 9; i++ {
		fmt.Scanf("%s", &s)
		p[i/3][i%3] = s
		if s == "x" {
			ix = i / 3
			iy = i % 3
		}
	}
	fmt.Println(p)
}
