package main

import (
	"fmt"
	"sort"
)

func dfs(s, cur string, u, start, state int, res *[]string) {
	if u == len(s) {
		*res = append(*res, cur)
		return
	}
	if u == 0 || s[u] != s[u-1] {
		start = 0
	}
	for i := start; i < len(s); i++ {
		if state>>i&1 == 0 {
			dfs(s, cur+string(s[i]), u+1, i+1, state+(1<<i), res)
		}
	}
}

func dfs1(s, cur string, u, state int, res *map[string]struct{}){
	if u == len(s){
		(*res)[cur]= struct{}{}
		return
	}
	for i:=0;i<len(s);i++{
		if ((state>>i)&1) == 0{
			dfs1(s, cur+string(s[i]), u+1, state+(1<<i), res)
		}
	}
}

func permutation(s string) []string {
	var res []string
	var tmp []rune
	for _, r := range []rune(s) {
		tmp = append(tmp, r)
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	var newS string
	for _, c := range tmp {
		newS += string(c)
	}
	//dfs(newS, "", 0, 0, 0, &res)
	ans := make(map[string]struct{})
	dfs1(newS, "", 0, 0, &ans)
	for k, _ :=range ans{
		res = append(res, k)
	}
	return res
}
func main() {
	fmt.Println(permutation("aab")) // want: ["aba","aab","baa"]
}
