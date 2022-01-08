package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	l, r := -10000.0, 10000.0
	for r-l > 1e-8 {
		mid := (r + l) / 2
		if mid*mid*mid >= x {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Printf("%.6f", r)
}
