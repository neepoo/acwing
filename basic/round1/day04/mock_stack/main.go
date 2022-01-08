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
	stk [N]int
	tt  = 0
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
			stk[tt] = x
		case "query":
			fmt.Println(stk[tt])
		case "pop":
			tt--
		case "empty":
			if tt == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
		m--
	}
}
