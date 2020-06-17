package dynamicprogramming

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积
// 你可以假设 n 不小于 2 且不大于 58
// https://leetcode-cn.com/problems/integer-break/description/ medium
func Test_IntegerBreak(t *testing.T) {
	n := 4
	t.Log(integerBreak(n))
}

// dp[n] = max(dp[n-1]*1, dp[n-2]*2 ... dp[1]*(n-1))
// 当n<=3时，dp[n]<n
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = common.Max(dp[i], j*common.Max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}
