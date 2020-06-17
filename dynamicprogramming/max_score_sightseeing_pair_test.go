package dynamicprogramming

import (
	"math"
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。
// 一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。
// 返回一对观光景点能取得的最高分。
// https://leetcode-cn.com/problems/best-sightseeing-pair
func Test_MaxScoreSightseeingPair(t *testing.T) {
	A := []int{8, 1, 5, 2, 6}
	t.Log(maxScoreSightseeingPair1(A))
}

// 超过时间限制
func maxScoreSightseeingPair(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	ret := math.MinInt32
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			ret = common.Max(ret, A[j]+A[i]+i-j)
		}
	}
	return ret
}

// 对于A[j]+A[i]+i-j来说，要取得最大即取A[j]-j和A[i]+i的和最大
// 将数组转换两个数组后取最大的两个元素相加即可(j>i)
func maxScoreSightseeingPair1(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	lowA := make([]int, len(A))
	uppA := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		uppA[i] = A[i] - i
		lowA[i] = A[i] + i
	}
	dp := make([]int, len(A)) // 前i个元素的最大值
	dp[0] = lowA[0]
	for i := 1; i < len(A); i++ {
		dp[i] = dp[i-1]
		if lowA[i] > dp[i-1] {
			dp[i] = lowA[i]
		}
	}
	ret := math.MinInt32
	for i := 1; i < len(A); i++ {
		if uppA[i]+dp[i-1] > ret {
			ret = uppA[i] + dp[i-1]
		}
	}
	return ret
}
