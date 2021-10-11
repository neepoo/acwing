package main

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestStr(t *testing.T) {
	ts := []struct {
		s, p string
		idx  int
	}{
		{
			"hello", "ll", 2,
		},
		{
			"aaaaa", "bba", -1,
		},
		{
			"", "", 0,
		},
	}
	for idx, i := range ts{
		t.Run("case"+strconv.Itoa(idx), func(t *testing.T) {
			got := strStr1(i.s, i.p)
			require.Equal(t, i.idx, got)
		})
	}
}
