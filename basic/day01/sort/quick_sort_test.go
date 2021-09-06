package sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		a []int
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "逆序排列",
			args: args{
				a: []int{5, 4, 3, 2, 1},
				l: 0,
				r: 4,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "两个元素",
			args: args{
				a: []int{2, 1},
				l: 0,
				r: 1,
			},
			want: []int{1, 2},
		},
		{
			name: "case3",
			args: args{
				a: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
				l: 0,
				r: 8,
			},
			want: []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
		},
		{
			name: "一个元素",
			args: args{
				a: []int{813},
				l: 0,
				r: 0,
			},
			want: []int{813},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.a, tt.args.l, tt.args.r)
			require.Equal(t, tt.want, tt.args.a)
		})
	}
}
