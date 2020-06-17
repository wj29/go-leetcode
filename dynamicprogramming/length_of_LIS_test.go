package dynamicprogramming

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 给定一个无序的整数数组，找到其中最长上升子序列的长度
// https://leetcode-cn.com/problems/longest-increasing-subsequence/ medium
func Test_LengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	t.Log(lengthOfLIS(nums))
}

// 最长子序列，f(n)的最长连续子序列一定是f(n-1)的最长连续子序列或者，加上nums[n]的后的最长子序列最大值大于nums[n]，那么f(n)则+1
// 以nums[i]为结束的最长上升子序列是之前所有子序列中最长上升子序列长度+1的最大值
// dp[i] = max(dp[i],dp[j]+1) for j in 0 to i
func lengthOfLIS(nums []int) int {
	ret := make([]int, len(nums))
	r := 0
	for i := range ret {
		ret[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				ret[i] = common.Max(ret[i], ret[j]+1)
			}
		}
		r = common.Max(r, ret[i])
	}
	return r
}
