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
	res, mod := div(s, b)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i])
	}
	fmt.Println()
	fmt.Println(mod)
}

func div(s string, b int) ([]int, int) {
	n1 := len(s)
	a := make([]int, 0, n1)
	for i := n1 - 1; i >= 0; i-- {
		a = append(a, int(s[i]-'0'))
	}
	res := make([]int, 0, n1)
	m := 0
	c := 0
	for i := n1 - 1; i >= 0; i-- {
		m = m*10 + a[i]
		c = m / b
		res = append(res, c)
		m %= b
	}
	for len(res) > 1 && res[0] == 0 {
		res = res[1:]
	}
	return res, m
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
	divRes, mod := a.DivMod(a, b, b)
	fmt.Println(divRes.String())
	fmt.Println(mod.String())
}
