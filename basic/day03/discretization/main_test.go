package main

import (
	"reflect"
	"testing"
)

func Test_unique(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				[]int{1, 2, 2, 3, 4, 4, 5, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
