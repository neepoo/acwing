package main

import (
	"fmt"
	"math/big"
)

func main() {
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
