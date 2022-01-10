package main

import (
	"fmt"
	"math/big"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	hand(s1, s2)
}

func hand(s1 string, s2 string) {
	n1, n2 := len(s1), len(s2)
	a, b := make([]int, n1), make([]int, n2)
	for i := n1 - 1; i >= 0; i-- {
		a[n1-1-i] = int(s1[i] - '0')
	}
	for i := n2 - 1; i >= 0; i-- {
		b[n2-1-i] = int(s2[i] - '0')
	}
	if cmp(a, b) {
		res := sub(a, b)
		for i := len(res) - 1; i >= 0; i-- {
			fmt.Print(res[i])
		}
	} else {
		res := sub(b, a)
		fmt.Print("-")
		for i := len(res) - 1; i >= 0; i-- {
			fmt.Print(res[i])
		}
	}
}

// a >= b
func sub(a []int, b []int) []int {
	res := make([]int, 0, len(a))
	c := 0
	for i := 0; i < len(a); i++ {
		c += a[i]
		if i < len(b) {
			c -= b[i]
		}
		res = append(res, (c+10)%10)
		if c < 0 {
			c = -1
		} else {
			c = 0
		}
	}
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	return res
}

func cmp(a []int, b []int) bool {
	n1 := len(a)
	n2 := len(b)
	if n1 != n2 {
		return n1 > n2
	}
	for i := n1 - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}

func stl(s1 string, s2 string) {
	var (
		a = new(big.Int)
		b = new(big.Int)
	)
	a, _ = a.SetString(s1, 10)
	b, _ = b.SetString(s2, 10)
	a.Sub(a, b)
	fmt.Println(a.String())
}
