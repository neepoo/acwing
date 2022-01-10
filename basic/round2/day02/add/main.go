package main

import (
	"fmt"
	"math/big"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	//stl(s1, s2)
	hand(s1, s2)
}

func hand(s1 string, s2 string) {
	n1, n2 := len(s1), len(s2)
	a := make([]int, n1)
	b := make([]int, n2)
	for i := n1 - 1; i >= 0; i-- {
		a[n1-1-i] = int(s1[i] - '0')
	}
	for i := n2 - 1; i >= 0; i-- {
		b[n2-1-i] = int(s2[i] - '0')
	}
	res := add(a, b)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
}

func add(a []int, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	i, c := 0, 0
	for i < len(a) || i < len(b) || c > 0 {
		if i < len(a) {
			c += a[i]
		}
		if i < len(b) {
			c += b[i]
		}
		res = append(res, c%10)
		c /= 10
		i++
	}
	return res
}

func stl(s1 string, s2 string) {
	var (
		a = new(big.Int)
		b = new(big.Int)
	)
	a, _ = a.SetString(s1, 10)
	b, _ = b.SetString(s2, 10)
	a.Add(a, b)
	fmt.Println(a.String())
}
