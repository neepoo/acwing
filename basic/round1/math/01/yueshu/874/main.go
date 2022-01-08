package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	n    int
	a, b int
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	buf := strings.Builder{}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		buf.WriteString(strconv.Itoa(gcd(a, b)) + "\n")
	}
	fmt.Print(buf.String())
}
