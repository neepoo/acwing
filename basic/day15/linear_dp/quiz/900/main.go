// https://www.acwing.com/problem/content/900/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 510

var (
	g [N][N]int
	f [N][N]int
	n int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// v表示上一层能到这个点的最大值, u表示第几层,从一开始
func dfs(u int) {
	if u == n+1 {
		return
	}
	for i := 1; i <= u; i++ {
		if i == 1 {
			f[u][i] = f[u-1][i] + g[u][i]
		} else if i == u {
			f[u][i] = f[u-1][i-1] + g[u][i]
		} else {
			f[u][i] = max(f[u-1][i], f[u-1][i-1]) + g[u][i]
		}
		dfs(u + 1)
	}
}

/*
10
-6
-4 -5
-3 7 5
3 7 -2 1
10 2 -6 2 -6
-8 3 8 6 7 9
-4 -10 0 -3 4 9 2
0 5 5 5 10 -6 -5 -4
-9 7 4 9 8 -5 -2 3 2
-7 -4 0 -10 -8 -4 3 -5 8 9


30
13
59 -63
-66 -50 47
0 20 -30 71
22 -12 26 -55 34
84 -32 80 93 86 49
-50 42 -43 49 -4 -16 -53
87 44 -26 9 76 -71 -47 -27
93 -2 66 -53 0 -50 -97 56 -38
29 52 42 -23 -3 54 69 84 35 -57
-30 -55 3 -13 -71 27 72 71 -99 -74 20
-86 -44 -86 58 -86 94 -76 -72 39 67 -59 95
-39 78 -65 -45 38 -83 -48 -56 -47 -48 66 93 -25
7 -99 22 -89 49 66 28 15 -66 -97 44 43 -83 59
-95 -90 82 -17 10 72 -49 -73 89 64 50 27 11 -52 -84
58 40 -50 92 -58 82 61 -95 12 50 97 -66 32 -34 74 -5
91 23 -56 16 -4 -23 6 85 74 66 13 -21 89 99 89 -73 -90
49 42 75 -81 -56 -34 -11 46 -31 -98 8 43 -13 -97 -7 -56 54 -72
68 28 -43 77 -40 -78 17 -73 58 -80 24 32 -64 -98 -83 70 9 -31 -40
-13 93 62 -46 31 70 -70 -63 70 71 -92 50 -97 25 24 -15 -84 86 52 -65
83 60 5 -65 -35 66 -31 -69 50 29 31 11 -98 69 99 -45 98 6 84 -2 -16
-42 -66 21 75 71 -85 -36 -2 9 59 78 -47 -44 -57 -4 77 -13 -26 50 -16 50 58
66 4 94 26 87 83 75 -38 61 -20 -68 24 -80 100 56 -84 15 -73 -67 -11 35 52 -63
-83 87 7 81 41 -55 -44 76 -99 14 -98 52 23 0 -78 -63 -45 31 79 -16 99 50 4 -35
12 64 -55 -23 30 -8 23 -93 80 33 -29 -68 91 -45 0 75 72 -34 77 99 -11 20 60 62 9
-39 45 81 -82 -58 -11 59 -84 -79 -78 -3 65 -100 91 -3 -18 -48 38 91 15 -18 18 24 -52 -7 -60
-25 68 -34 -3 -7 5 80 30 54 24 10 15 11 0 61 -28 -25 53 96 1 21 -31 51 -80 -4 -87 14
2 50 -71 55 -51 75 55 -14 -12 -92 98 -51 -56 -34 -47 5 -54 -98 -35 -22 -79 -88 81 -35 92 -16 65 10
94 -16 26 46 -33 -29 100 -65 51 2 -32 -13 -3 41 99 -11 85 -17 67 21 49 -49 78 -45 -84 84 -42 -96 50
-93 89 -49 -74 -100 44 -21 -71 51 5 84 -6 -10 -59 -68 97 92 41 -4 79 67 39 -82 74 -31 -87 68 -2 -4 63
*/

func solution() {
	for u := 1; u <= n; u++ {
		for i := 1; i <= u; i++ {
			if i == 1 {
				f[u][i] = f[u-1][i] + g[u][i]
			} else if i == u {
				f[u][i] = f[u-1][i-1] + g[u][i]
			} else {
				f[u][i] = max(f[u-1][i], f[u-1][i-1]) + g[u][i]
			}
		}
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, N*N)
	Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			var a int
			if j == i {
				Fscanf(reader, "%d\n", &a)
			} else {
				Fscanf(reader, "%d ", &a)
			}
			g[i][j] = a
			f[i][j] = a
		}
	}
	//dfs(1)
	solution()
	ans := -0x3f3f3f3f
	for i := 1; i <= n; i++ {
		ans = max(ans, f[n][i])
	}
	Println(ans)
}
