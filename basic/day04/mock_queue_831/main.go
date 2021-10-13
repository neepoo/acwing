package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010

var (
	q  [N]int
	tt = -1
	hh = 0
)

func main() {
	var m int
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, N*6)
	fmt.Fscanf(reader, "%d\n", &m)
	scanner := bufio.NewScanner(reader)
	for m > 0 && scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "push":
			x, _ := strconv.Atoi(tokens[1])
			tt++
			q[tt] = x
		case "query":
			fmt.Println(q[hh])
		case "pop":
			hh++
		case "empty":
			if hh > tt {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
		m--
	}
}
