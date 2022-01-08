package dichotomous

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumRange(t *testing.T) {
	testCase := []struct {
		nums []int
		x    int
		l    int
		r    int
	}{
		{
			[]int{1, 2, 2, 3, 3, 4}, 3, 3, 4,
		},
		{
			[]int{1, 2, 2, 3, 3, 4}, 4, 5, 5,
		},
		{
			[]int{1, 2, 2, 3, 3, 4}, 5, -1, -1,
		},
	}

	for idx, ts := range testCase {
		t.Run("test case "+strconv.Itoa(idx), func(t *testing.T) {
			gotL, gotR := numRange(ts.nums, ts.x)
			require.Equal(t, ts.l, gotL, "左边界不相等")
			require.Equal(t, ts.r, gotR, "右边界不相等")
		})
	}
}
