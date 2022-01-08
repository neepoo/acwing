package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func getAscii(a rune) uint64 {
	return uint64(a - '0' + 1)
}

func rabinKarp(t, s string) []int {
	var P uint64 = 131
	p := make([]uint64, 1000010)
	h := make([]uint64, 1000010)
	p[0] = 1
	for i := 0; i < len(t); i++ {
		p[i+1] = p[i] * P
		h[i+1] = h[i]*P + getAscii(rune(t[i]))
	}
	var haseS uint64
	for i := 0; i < len(s); i++ {
		haseS = haseS*P + getAscii(rune(s[i]))
	}
	var res []int
	for i := 0; i+len(s) <= len(t); i++ {
		currentHash := h[i+len(s)] - h[i]*p[len(s)]
		if currentHash == haseS {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	var n, m int
	var s, t string
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 1e6+1e5+100)
	fmt.Fscanf(reader, "%d\n", &n)
	fmt.Fscanf(reader, "%s\n", &s)
	fmt.Fscanf(reader, "%d\n", &m)
	fmt.Fscanf(reader, "%s\n", &t)
	res := rabinKarp(t, s)
	buf := bytes.NewBufferString("")

	for _, i := range res {
		buf.WriteString(strconv.Itoa(i) + " ")
	}
	fmt.Println(buf.String())
}
