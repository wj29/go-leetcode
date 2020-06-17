package dynamicprogramming

import (
	"testing"
)

// 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，
// 你都可以从 + 或 -中选择一个符号添加在前面
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数
// https://leetcode-cn.com/problems/target-sum  medium
func Test_FindTargetSumWays(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	s := 3
	t.Log(findTargetSumWays(nums, s))
}

// dp[i][j]表示从0-i个元素经过添加符号后得到数值j的方法数
// 那么对于元素i来说，添加元素i，可以加上或减去 // 思路不对，无法实现

// 问题变化后，找到一个正子集P和一个负子集N使得sum(P)-sum(N)=target
// 经过数学变换 sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
// 2 * sum(P) = target + sum(nums)
// 即找到一个子集等于数组总和+target的和的一半
// 此时问题转换成can_partition_test.go，又不同于，因为此题必须选择第i个元素
// dp[i]代表的含义是从nums中取数相加和为i时有多少种取法
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum < S {
		return 0
	}
	sum += S
	if sum%2 != 0 {
		return 0
	}
	dp := make([]int, sum/2+1)
	dp[0] = 1 // 从元素中一个都不取的方式使其和为0
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= nums[i]; j-- { //
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}
	return dp[sum/2]
}
