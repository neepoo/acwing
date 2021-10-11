package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "case1", args{s: "III"}, 3},
		{ "case2", args{s: "IV"}, 4},
		{ "case3", args{s: "IX"}, 9},
		{ "case4", args{s: "LVIII"}, 58},
		{ "case5", args{s: "MCMXCIV"}, 1994},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
