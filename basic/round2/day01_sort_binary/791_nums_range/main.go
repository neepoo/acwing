package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1e5

var (
	n, q int
	nums [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscan(reader, &n, &q)
	for i := 0; i < n; i++ {
		Fscan(reader, &nums[i])
	}
	for i := 0; i < q; i++ {
		var x int
		Fscan(reader, &x)

		l, r := 0, n-1
		for l < r {
			mid := (l + r) / 2
			if nums[mid] >= x {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if nums[l] != x {
			Println(-1, -1)
		} else {
			Print(l, " ")
			l, r = 0, n-1
			for l < r {
				mid := (l + r + 1) / 2
				if nums[mid] <= x {
					l = mid
				} else {
					r = mid - 1
				}
			}
			Println(r)
		}
	}
}
