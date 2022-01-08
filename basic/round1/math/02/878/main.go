package main

import (
	"bufio"
	"fmt"
	"os"
)

var n uint64

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var b, m uint64
	var i uint64
	for ; i < n; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &b, &m)
		if b%m == 0 {
			fmt.Println("impossible")
		} else {
			fmt.Println(qmi(b, m-2, m))
		}
	}
}

func qmi(a uint64, k uint64, p uint64) uint64 {
	var res uint64 = 1
	// 将k表示成二进制, 哪一位是1,就乘上 pow(a, pow(2, i))
	for k > 0 {
		if k&1 == 1 {
			res = res * a % p
		}
		k >>= 1
		a = a * a % p
	}
	return res
}
