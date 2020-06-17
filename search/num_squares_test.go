package search

import (
	"testing"
)

func Test_NumSquares(t *testing.T) {
	n := 12
	t.Log(numSquares(n))
}

// 不会
func numSquares(n int) int {
	return 0
}
