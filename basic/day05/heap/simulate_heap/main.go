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
	n, size int
	q       [N]int
	ph, hp  [N]int
)

func main() {
	m:=0
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d\n", &n)
	scanner := bufio.NewScanner(reader)
	for n > 0 && scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		op := tokens[0]
		switch op {
		case "I":
			x, _ := strconv.Atoi(tokens[1])
			size++
			m++
			ph[m]=size
			hp[size] = m
			q[size] = x
			up(size)
		case "PM":
			fmt.Println(q[1])
		case "DM":
			swapHeap(1, size)
			size--
			down(1)
		case "D":
			k, _ := strconv.Atoi(tokens[1])
			k = ph[k]
			swapHeap(k, size)
			size--
			down(k)
			up(k)
		case "C":
			k, _ := strconv.Atoi(tokens[1])
			k = ph[k]
			x, _ := strconv.Atoi(tokens[2])
			q[k] = x
			down(k)
			up(k)
		}
		n--
	}
}

func swap(nums []int, a, b int){
	nums[a], nums[b] = nums[b], nums[a]
}

func swapHeap(a, b int) {
	swap(ph[:], hp[a], hp[b])
	swap(hp[:], a, b)
	swap(q[:], a, b)
}

func down(i int) {
	u := i
	if 2*i <= size && q[2*i] < q[u] {
		u = 2 * i
	}
	if 2*i+1 <= size && q[2*i+1] < q[u] {
		u = 2*i + 1
	}
	if u != i {
		swapHeap(i, u)
		down(u)
	}
}

func up(i int) {
	//for {
	//	j := i / 2
	//	if j == 0 || q[j] > q[i]{
	//		break
	//	}
	//	swapHeap(i, j)
	//	i = j
	//}
	for i/2 > 0 && q[i/2] > q[i] {
		swapHeap(i, i/2)
		i /= 2
	}
}
