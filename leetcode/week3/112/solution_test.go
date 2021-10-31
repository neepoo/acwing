package leetcode

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case114",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -2,
						Left: &TreeNode{
							Val: 1,
							Left: &TreeNode{
								Val:   -1,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: -3,
						Left: &TreeNode{
							Val:   -2,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
				targetSum: -1,
			},
			want: true,
		},
		{
			name: "Case2",
			args: args{
				root:      &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				targetSum: 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
