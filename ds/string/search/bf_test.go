package search

import (
	"strings"
	"testing"
)

func Test_bs1(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "",
				p: "",
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				s: "",
				p: "11",
			},
			want: -1,
		},
		{
			name: "case3",
			args: args{
				s: "abcdef",
				p: "de",
			},
			want: strings.Index("abcdef", "de"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bs1(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("bs1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bs2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "",
				p: "",
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				s: "",
				p: "11",
			},
			want: -1,
		},
		{
			name: "case3",
			args: args{
				s: "abcdecdef",
				p: "def",
			},
			want: strings.Index("abcdecdef", "def"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bs2(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("bs2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmp(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				s: "abaabaabeca",
				p: "abaabe",
			},
			want: strings.Index("abaabaabeca", "abaabe"),
		},
		{
			name: "Case2",
			args: args{
				s: "sdfwegfdhabc234rabcdabcabcd",
				p: "abcabcd",
			},
			want: strings.Index("sdfwegfdhabc234rabcdabcabcd", "abcabcd"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kmp(tt.args.s, tt.args.p)
			if got != tt.want {
				t.Fatal("kmp error")
			}
		})
	}
}
