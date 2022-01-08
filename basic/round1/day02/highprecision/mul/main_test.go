package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSub(t *testing.T) {
	testCase := []struct {
		a    string
		b    int
		want string
		name string
	}{
		{
			a:    "22",
			b:    3,
			want: "66",
			name: "case1",
		},
		{
			a:    "22",
			b:    100,
			want: "2200",
			name: "case2",
		},
		{
			a:    "12345",
			b:    0,
			want: "0",
			name: "case3",
		},
	}

	for _, ts := range testCase {
		t.Run(ts.name, func(t *testing.T) {
			got := Mul(ts.a, ts.b)
			require.Equal(t, ts.want, got)
		})
	}
}
