package acwing

import (
	"reflect"
	"testing"
)

func Test_deleteDuplication(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Case1",
			args: args{
				head: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "Case2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val:  3,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{Val: 2, Next: &ListNode{
				Val:  3,
				Next: nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplication(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
