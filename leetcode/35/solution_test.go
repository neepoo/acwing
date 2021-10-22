package solution

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1", args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			2,
		},
		{
			"case2", args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			1,
		},
		{
			"case3", args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			4,
		},
		{
			"case4", args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			0,
		},
		{
			"case5", args{
				nums:   []int{1},
				target: 0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
