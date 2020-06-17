package dynamicprogramming

import (
	"math"
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点
// https://leetcode-cn.com/problems/triangle  medium
func Test_MinimumTotal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	//triangle := [][]int{{2}, {3, 4}}

	t.Log(minimumTotal(triangle))
}

// 对于dp[i][j]的值一定是min(dp[i-1][j],dp[i-1][j-1])+nums[i][j]，它只可能由上述两条路径得到
// 比较最后一行的dp[i][j]值即可
func minimumTotal(triangle [][]int) int {
	if len(triangle) <= 0 {
		panic("invalid length")
	}
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle))
	}
	dp[0][0] = triangle[0][0]
	ret := math.MaxInt32

	for i := 1; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j] // 没有前一位
			} else if i == j {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j] // 没有上一位
			} else {
				dp[i][j] = common.Min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	for _, v := range dp[len(triangle)-1] {
		if v < ret {
			ret = v
		}
	}
	return ret
}
