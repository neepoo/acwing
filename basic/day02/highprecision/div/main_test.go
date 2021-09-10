package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiv(t *testing.T) {
	testCase := []struct {
		a     string
		b     int
		shang string
		yushu int
		name  string
	}{
		{
			a:     "7",
			b:     2,
			shang: "3",
			yushu: 1,
			name:  "case1",
		},
		{
			a:     "1000",
			b:     10,
			shang: "100",
			yushu: 0,
			name:  "case2",
		},
	}

	for _, ts := range testCase {
		t.Run(ts.name, func(t *testing.T) {
			gotShang, gotYuSHu := Div(ts.a, ts.b)
			require.Equal(t, ts.shang, gotShang)
			require.Equal(t, ts.yushu, gotYuSHu)
		})
	}
}
