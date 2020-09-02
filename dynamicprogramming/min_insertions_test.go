package dynamicprogramming

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 1312. 让字符串成为回文串的最少插入次数
// https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
func Test_MinInsertions(t *testing.T) {
	s := "ab"
	t.Log(minInsertions(s))
}

func minInsertions(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = common.Min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][len(s)-1]
}
