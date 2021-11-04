package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_set(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Case1",
			args: args{s: "12345678"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set(tt.args.s)
			require.Equal(t, [2][4]int{
				{1, 2, 3, 4},
				{8, 7, 6, 5},
			}, g)
		})
	}
}

func Test_get(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Case1",
			want: "12345678",
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g = [2][4]int{
				{1, 2, 3, 4},
				{8, 7, 6, 5},
			}
			require.Equal(t, tt.want, get())
		})
	}
}