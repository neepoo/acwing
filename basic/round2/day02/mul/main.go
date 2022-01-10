package main

import (
	"fmt"
	"math/big"
)

func main() {
	//stl()
	handWriting()
}

func handWriting() {
	var s string
	var b int
	fmt.Scan(&s)
	fmt.Scan(&b)
	res := mul(s, b)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
	}
}

func mul(s string, b int) []int {
	n1 := len(s)
	a := make([]int, 0, n1)
	for i := n1 - 1; i >= 0; i-- {
		a = append(a, int(s[i]-'0'))
	}
	res := make([]int, 0, n1+b/10)
	c := 0
	for i := 0; i < n1 || c > 0; i++ {
		if i < n1 {
			c += b * a[i]
		}
		res = append(res, c%10)
		c /= 10
	}
	// delete leading zero
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	return res
}

func stl() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	var (
		a = new(big.Int)
		b = new(big.Int)
	)
	a, _ = a.SetString(s1, 10)
	b, _ = b.SetString(s2, 10)
	a.Mul(a, b)
	fmt.Println(a.String())
}
