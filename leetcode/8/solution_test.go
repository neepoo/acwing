package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "Case2",
			args: args{s: "91283472332"},
			want: 2147483647,
		},
		{
			name: "Case3",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "Case4",
			args: args{s: "234"},
			want: 234,
		},
		{
			name: "Case5",
			args: args{s: "-901"},
			want: -901,
		},
		{
			name: "Case6",
			args: args{s: "+1"},
			want: 1,
		},
		{
			name: "Case7",
			args: args{s: "9223372036854775808"},
			want: 2147483647,
		},
		{
			name: "Case8",
			args: args{s: "2147483646"},
			want: 2147483646,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{s: "42"},
			want: 42,
		},
		{
			name: "Case2",
			args: args{s: "   -42"},
			want: -42,
		},
		{
			name: "Case3",
			args: args{s: "4193 with words"},
			want: 4193,
		}, {
			name: "Case4",
			args: args{s: "words and 987"},
			want: 0,
		},
		{
			name: "Case5",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "Case6",
			args: args{s: "+-222"},
			want: 0,
		},
		{
			name: "Case7",
			args: args{s: "+"},
			want: 0,
		},
		{
			name: "Case8",
			args: args{s: "   "},
			want: 0,
		},
		{
			name: "Case9",
			args: args{s: "9223372036854775808"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
