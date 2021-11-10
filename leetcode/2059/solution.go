package leetcode

type node struct {
	val int // 值
	d   int // start 变更到val经过的操作次数
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

func xor(a, b int) int {
	return a ^ b
}

func minimumOperations(nums []int, start int, goal int) int {
	m := make(map[int]bool)
	var q [1010]node
	hh, tt := 0, -1
	tt++
	q[tt] = node{start, 0}
	m[start] = true
	op := []func(a, b int) int{add, sub, xor}
	for hh <= tt {
		head := q[hh]
		hh++
		for i := 0; i < len(nums); i++ {
			for j := 0; j < 3; j++ {
				x := op[j](head.val, nums[i])
				// 应该时计算状态的时候就判断是否达到了目标状态,如果是目标状态返回就可以
				// 而不应该等到出队列的时候判断
				if x == goal{
					return head.d + 1
				}
				if 0 <= x && x <= 1000 && !m[x]{
					m[x] = true
					tt++
					q[tt] = node{
						val: x,
						d:   head.d + 1,
					}
				}
			}
		}
	}
	return -1
}
