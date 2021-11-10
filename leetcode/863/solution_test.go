package leetcode

import (
	"reflect"
	"testing"
)

func Test_distanceK(t *testing.T) {
	target := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
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
	}
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{
				root: &TreeNode{
					Val:  3,
					Left: target,
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   0,
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
				k: 2,
			},
			want: []int{1, 7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceK(tt.args.root, target, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distanceK1(t *testing.T) {
	target := &TreeNode{
		Val: 2,
	}
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: target,
					},
					Right: nil,
				},
				k: 1,
			},
			want: []int{1},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceK(tt.args.root, target, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}
