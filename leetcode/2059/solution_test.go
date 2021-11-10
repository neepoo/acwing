package leetcode

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums  []int
		start int
		goal  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				nums:  []int{1, 3},
				start: 6,
				goal:  4,
			},
			want: 2,
		},
		{
			name: "Case2",
			args: args{
				nums:  []int{2, 4, 12},
				start: 2,
				goal:  12,
			},
			want: 2,
		},
		{
			name: "Case3",
			args: args{
				nums:  []int{3, 5, 7},
				start: 0,
				goal:  -4,
			},
			want: 2,
		},
		{
			name: "Case4",
			args: args{
				nums:  []int{2, 8, 16},
				start: 0,
				goal:  1,
			},
			want: -1,
		},
		{
			name: "Case5",
			args: args{
				nums:  []int{1},
				start: 0,
				goal:  3,
			},
			want: 3,
		},
		{
			name: "Case6",
			args: args{
				nums:  []int{-21, 36, -12, 43, -4, -52, -93, 5, 12, 81, -90, 7, -31, -97, -49, 93, -65, 82, -37, 29, 87, -36, 70, 51, 60, -19, -73, -32, -13, -51, -23, 50},
				start: 4,
				goal:  789,
			},
			want: 9,
		},
	}
	/*
		[-21,36,-12,43,-4,-52,-93,5,12,81,-90,7,-31,-97,-49,93,-65,82,-37,29,87,-36,70,51,60,-19,-73,-32,-13,-51,-23,50]
		4
		789
	*/
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums, tt.args.start, tt.args.goal); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
