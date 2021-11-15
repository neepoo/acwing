package main

//
//func cuttingRope(n int) int {
//	var f [1010]int
//	// 枚举绳子长度
//	f[1] = 1
//	f[2] = 1
//	for i := 1; i <= n; i++ {
//		// 枚举左边绳子的长度
//		for j := 1; j < i; j++ {
//			f[i] = max(f[i], j * (i-j))
//			f[i] = max(f[i], j*f[i-j])
//		}
//	}
//	return f[n] % 1000000007
//}
const mod = 1000000007

func qmi(a, b, c int) int {
	var res = 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a %c
		}
		b >>= 1
		a = a * a % c
	}
	return res
}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	switch b {
	case 0:
		return qmi(3, a, mod) % mod
	case 1:
		return qmi(3, a-1, mod) * 4 %mod
	default:
		return qmi(3, a, mod) * 2 % mod
	}
}
func main() {
	//println(cuttingRope(10))
	println(cuttingRope(120))
	println(qmi(3, 3, mod))
}
