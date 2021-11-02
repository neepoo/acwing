package main

import "fmt"

var res []string
var m = map[string]string{
	"2":"abc",
	"3":"def",
	"4":"ghi",
	"5":"jkl",
	"6":"mno",
	"7":"pqrs",
	"8":"tuv",
	"9":"wxyz",

}

func letterCombinations(str string) []string {
	if len(str) == 0{
		return res
	}
	dfs(str, 0, "")
	return res
}

func dfs(str string, u int, path string){
	if u == len(str) {
		res = append(res, path)
	}else{
		for _, ch := range m[string(str[u])] {
			dfs(str, u+1, path+string(ch))
		}
	}
}

func main() {
	letterCombinations("")
	fmt.Println(res)
}
