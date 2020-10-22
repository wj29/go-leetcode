package dynamicprogramming

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 486. 预测赢家
// https://leetcode-cn.com/problems/predict-the-winner/
func Test_PredictTheWinner(t *testing.T) {
	nums := []int{1, 5, 233, 7}
	nums = []int{1, 4, 2}
	t.Log(PredictTheWinner(nums))
}

// dp[i][j]表示从数组nums i-j中两个玩家的分数差
// 对于nums i-j，玩家1有两个选择，
// 当玩家1选择了nums[i]时，那么dp[i-1][j]表示从i-1到j的最优解分数(玩家2的)，那么dp[i][j] = nums[i]-dp[i+1][j]
// 当玩家1选择了nums[j]时，同理dp[i][j] = nums[j]-dp[i][j-1]
// 那么对于dp[i][j]我们需要先知道dp[i+1][j]和dp[i][j-1]的值，即遍历顺序为从下往上，从左向右
// base: i == j 时，dp[i][j] = nums[i]，i == j-1时，dp[i][j] = abs(nums[i]-nums[j])
func PredictTheWinner(nums []int) bool {
	// 对于偶数个的数组，玩家1一定获胜
	// 如果玩家1选择拿法A，玩家2选择拿法B，玩家1输了。则玩家1换一种拿法选择拿法B，因为玩家1是先手，所以玩家1一定可以获胜
	if len(nums)%2 == 0 {
		return true
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if j == i+1 {
				dp[i][j] = common.Abs(nums[i], nums[j])
			} else {
				dp[i][j] = common.Max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
			}
		}
	}
	return dp[0][len(nums)-1] >= 0
}
