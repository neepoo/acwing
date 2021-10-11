package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name:"case1",args: args{strs: []string{"flower", "flow", "flight"}}, want: "fl"},
		{name:"2",args: args{strs: []string{"dog", "racecar", "car"}}, want: ""},
		{name:"3",args: args{strs: []string{"cir", "car"}}, want: "c"},
		{name:"4",args: args{strs: []string{"flower","flower","flower","flower"}}, want: "flower"},
		{name:"5",args: args{strs: []string{"babb", "caa"}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
