/*
https://www.acwing.com/problem/content/1012/

某国为了防御敌国的导弹袭击，发展出一种导弹拦截系统。

但是这种导弹拦截系统有一个缺陷：虽然它的第一发炮弹能够到达任意的高度，但是以后每一发炮弹都不能高于前一发的高度。

某天，雷达捕捉到敌国的导弹来袭。

由于该系统还在试用阶段，所以只有一套系统，因此有可能不能拦截所有的导弹。

输入导弹依次飞来的高度（雷达给出的高度数据是不大于30000的正整数，导弹数不超过1000），计算这套系统最多能拦截多少导弹，如果要拦截所有导弹最少要配备多少套这种导弹拦截系统。

输入格式
共一行，输入导弹依次飞来的高度。

输出格式
第一行包含一个整数，表示最多能拦截的导弹数。

第二行包含一个整数，表示要拦截所有导弹最少要配备的系统数。

数据范围
雷达给出的高度数据是不大于 30000 的正整数，导弹数不超过 1000。

输入样例：
389 207 155 300 299 170 158 65
输出样例：
6
2
*/
package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1010

var (
	n       int
	a, f, g [N]int
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")
	n = len(tokens)
	var ans int
	for idx, s := range tokens {
		v, _ := strconv.Atoi(s)
		a[idx+1] = v
	}
	for i := 1; i <= n; i++ {
		f[i] = 1
		for j := 1; j < i; j++ {
			if a[j] >= a[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	Println(ans)

	// question2
	var cnt int
	for i := 1; i <= n; i++ {
		k := 1 // 第几个导弹系统
		// 找到比他高或相等的最低的一个
		for k <= cnt && a[i] > g[k] {
			k++
		}
		g[k] = a[i]
		if k > cnt{
			cnt = k
		}
	}
	Println(cnt)
}
