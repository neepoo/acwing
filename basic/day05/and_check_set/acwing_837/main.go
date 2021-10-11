package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.acwing.com/problem/content/839/

const N = 100010

var (
	n, m    int
	p, size [N]int
)

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 6*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 1; i <= n; i++ {
		p[i] = i
		size[i] = 1
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() && m > 0 {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		op := tokens[0]

		switch op {
		case "C":
			a, _ := strconv.Atoi(tokens[1])
			b, _ := strconv.Atoi(tokens[2])
			if find(a) == find(b) {
				continue
			}
			size[find(b)] += size[find(a)]
			p[find(a)] = find(b)
		case "Q1":
			a, _ := strconv.Atoi(tokens[1])
			b, _ := strconv.Atoi(tokens[2])
			if find(a) == find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		case "Q2":
			a, _ := strconv.Atoi(tokens[1])
			fmt.Println(size[find(a)])
		}
		m--
	}
}
