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
	n, size /*表示当前堆中的元素个数*/, m/*表示第几个插入的数*/ int
	q [N]int
	ph /*ph[i]=j表示第i个插入的数在堆中的下标是j*/, hp/*hp[j]=i表示堆中下标是j的数是第i个插入的*/ [N]int
)

func main() {
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
			// 第m个插入的
			ph[m] = size
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

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

// a, b 表示堆中节点的下标,从1开始
func swapHeap(a, b int) {
	// 交换操作是幂等, 即swap(a, b)和swap(b, a)产生的效果一样,因此下面两行谁在前面都无所谓
	swap(ph[:], hp[a], hp[b]) // hp[a],hp[b] 分别表示堆中的a,b这两个节点是第几个插入的
	swap(hp[:], a, b)         // 因为hp数组中的下标表示的就是堆中元素的下标,因此直接swap(a,b)就可以
	swap(q[:], a, b)          // 交换堆中两个节点的值
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
	for i/2 > 0 && q[i/2] > q[i] {
		swapHeap(i, i/2)
		i /= 2
	}
}
