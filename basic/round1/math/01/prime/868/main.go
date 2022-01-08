// 试除法判定素数
package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &x)
		if isPrime(x) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
