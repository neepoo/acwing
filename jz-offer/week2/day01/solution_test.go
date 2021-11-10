package day01

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				m: 2,
				n: 3,
				k: 1,
			},
			want: 3,
		},

		{
			name: "Case2",
			args: args{
				m: 3,
				n: 1,
				k: 0,
			},
			want: 1,
		},
		{
			name: "Case3",
			args: args{
				m: 3,
				n: 2,
				k: 17,
			},
			want: 6,
		},
		{
			name: "Case4",
			args: args{
				m: 4,
				n: 5,
				k: 7,
			},
			want: 20,
		},
		{
			name: "Case5",
			args: args{
				m: 40,
				n: 40,
				k: 18,
			},
			want: 1484,
		},
		{
			name: "Case6",
			args: args{
				m: 0,
				n: 5,
				k: 6,
			},
			want: 1,
		},
		{
			name: "Case7",
			args: args{
				m: 3,
				n: 10,
				k: 10,
			},
			want: 29,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.k, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
