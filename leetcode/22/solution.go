package leetcode


func generateParenthesis(n int) []string {
	var res []string
	dfs(n, 0, 0, "", &res)
	return res
}

func dfs(n, lc, rc int, seq string, res *[]string) {
	if lc == n && rc == n{
		*res = append(*res, seq)
		return
	}else{
		if lc < n {
			dfs(n, lc+1, rc, seq+"(", res)
		}
		if rc < n && lc > rc {
			dfs(n, lc, rc+1, seq+")", res)
		}
	}
}