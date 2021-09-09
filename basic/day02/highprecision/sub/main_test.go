package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSub(t *testing.T) {
	testCase := []struct {
		a    string
		b    string
		name string
		want string
	}{
		{
			a:    "5555",
			b:    "5555",
			name: "case5",
			want: "0",
		},
		{
			a:    "6352309719",
			b:    "4523128104",
			name: "case5",
			want: "1829181615",
		},
		{
			a:    "100",
			b:    "0",
			name: "case1",
			want: "100",
		},
		{
			a:    "987654321",
			b:    "876543210",
			name: "case2",
			want: "111111111",
		},
		{
			a:    "876543210",
			b:    "987654321",
			name: "case3",
			want: "-111111111",
		},
		{
			a:    "32",
			b:    "11",
			name: "case4",
			want: "21",
		},

	}
	for _, ts := range testCase {
		t.Run(ts.name, func(t *testing.T) {
			got := sub(ts.a, ts.b)
			require.Equal(t, ts.want, got)
		})
	}
}
