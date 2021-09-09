package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	testCase := []struct {
		Name string
		a    string
		b    string
		want string
	}{
		{
			Name: "case error",
			a:    "2497221268",
			b:    "6022247976",
			want: "8519469244",
		},
		{
			Name: "case5",
			a:    "9999",
			b:    "1",
			want: "10000",
		},
		{
			Name: "case6",
			a:    "1",
			b:    "999999",
			want: "1000000",
		},
		{
			Name: "case1",
			a:    "22",
			b:    "33",
			want: "55",
		},
		{
			Name: "case2",
			a:    "0",
			b:    "9",
			want: "9",
		},
		{
			Name: "case3",
			a:    "3333",
			b:    "22",
			want: "3355",
		},
		{
			Name: "case4",
			a:    "1",
			b:    "9",
			want: "10",
		},
	}

	for _, ts := range testCase {
		got := Add(ts.a, ts.b)
		require.Equal(t, ts.want, got)
	}
}
