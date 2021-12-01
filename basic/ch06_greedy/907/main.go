package main
// https://www.acwing.com/problem/content/907/
import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

const N = 100010

type pair struct {
	x, y int
}

var (
	n   int
	q   [N]pair
	res int
	st  = -0x3f3f3f3f
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscan(reader, &n)
	for i := 0; i < n; i++ {
		Fscan(reader, &q[i].x, &q[i].y)
	}
	sort.Slice(q[:n], func(i, j int) bool {
		return q[i].y < q[j].y
	})
	for _, rg := range q[:n] {
		if st < rg.x || st > rg.y {
			res++
			st = rg.y
		}
	}
	Println(res)
}
