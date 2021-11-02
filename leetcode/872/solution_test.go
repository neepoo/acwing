package leetcode

import "testing"

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case1",
			args: args{
				root1: &TreeNode{
					Val: 1,
					Left: &TreeNode{
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
				root2: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: false,
		},
		{
			name: "Case2",
			args: args{
				root1: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   5,
						Left:  &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{
								Val:   7,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   4,
								Left:  nil,
								Right: nil,
							},
						},
					},
					Right: &TreeNode{
						Val:   1,
						Left:  &TreeNode{
							Val:   9,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   8,
							Left:  nil,
							Right: nil,
						},
					},
				},
				root2: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   5,
						Left:  &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   1,
						Left:  &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  &TreeNode{
								Val:   9,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   8,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
