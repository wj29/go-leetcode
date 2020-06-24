package dynamicprogramming

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 数组的每个索引作为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。
// 每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。
// 您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。
// https://leetcode-cn.com/problems/min-cost-climbing-stairs easy
func Test_MinCostClimbingStairs(t *testing.T) {
	cost := []int{10, 15, 20}
	t.Log(minCostClimbingStairs(cost))
}

// 对于最后一阶楼梯，可以通过前一阶或者前二阶爬上
// dp[i]表示爬上第i个阶梯需要的最少花费，那么dp方程为
// 构造最后一位台阶，花费为0
// dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i])
func minCostClimbingStairs(cost []int) int { // len(cost) >= 2
	if len(cost) == 2 {
		return common.Min(cost[0], cost[1])
	}
	cost = append(cost, 0)
	dp := make([]int, len(cost))
	// 初始化base
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(dp); i++ {
		dp[i] = common.Min(dp[i-1]+cost[i], dp[i-2]+cost[i])
	}
	return dp[len(cost)-1]
}
