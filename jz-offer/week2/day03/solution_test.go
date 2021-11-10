package day03

import "testing"

func TestNumberOf1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{n: 9},
			want: 2,
		},
		{
			name: "Case2",
			args: args{n: -2},
			want: 31,
		},
		{
			name: "Case3",
			args: args{n: -1},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOf1(tt.args.n); got != tt.want {
				t.Errorf("NumberOf1() = %v, want %v", got, tt.want)
			}
		})
	}
}
