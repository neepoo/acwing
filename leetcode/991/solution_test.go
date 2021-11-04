package leetcode

import "testing"

func Test_brokenCalc(t *testing.T) {
	type args struct {
		s int
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				s: 1024,
				t: 1,
			},
			want: 1023,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := brokenCalc(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("brokenCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
