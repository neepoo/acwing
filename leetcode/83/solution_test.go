package solution

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val:  3,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.args.head)
			require.Equal(t, tt.want, got)
		})
	}
}
