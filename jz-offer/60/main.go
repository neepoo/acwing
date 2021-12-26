package main

import "fmt"

func mi(a, x int) float64 {
	var res = 1
	for i := 1; i <= x; i++ {
		res *= a
	}
	return float64(res)
}

func dicesProbability(n int) []float64 {
	const (
		N = 12
		M = N * 6
	)
	var f [N][M]float64
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= 6*i; j++ {
			for k := 1; k <= min(6, j); k++ {
				f[i][j] += f[i-1][j-k]

			}
		}
	}
	var res []float64
	for i := n; i <= 6*n; i++ {
		res = append(res, f[n][i]/mi(6, n))
	}
	return res
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
func main() {
	fmt.Println(dicesProbability(2))
}
