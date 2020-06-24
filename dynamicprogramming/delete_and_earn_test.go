package dynamicprogramming

import (
	"github.com/wujiang1994/go-leetcode/common"
	"math"
	"testing"
)

// 给定一个整数数组 nums ，你可以对它进行一些操作。
// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除每个等于 nums[i] - 1 或 nums[i] + 1 的元素。
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
// https://leetcode-cn.com/problems/delete-and-earn medium
func Test_DeleteAndEarn(t *testing.T) {
	nums := []int{2, 2, 3, 3, 3, 4}
	t.Log(deleteAndEarn(nums))
}

// 以nums[i]的值作为下标，相同的的nums[i]的总和作为值构建新的数组，转换成打家劫舍问题
// 假若dp[i-1]取值不包括nums[i-1]，那么dp[i]的值一定是dp[i-1]+nums[i]，没有取nums[i-1]时dp[i-1]=dp[i-2]，
// 假若dp[i-1]取值包括了nums[i-1]，那么dp[i]的值不能取nums[i]，那么dp[i] = max(dp[i-1], dp[i-2]+nums[i])
// 那么dp[i] = max(dp[i-1], dp[i-2]+nums[i])
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, m := math.MinInt32, make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]] += nums[i]
		} else {
			m[nums[i]] = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	newNums := make([]int, max+1)
	for k, v := range m {
		newNums[k] = v
	}
	dp := make([]int, len(newNums)+1)
	dp[0], dp[1] = 0, newNums[0]
	for i := 2; i < len(dp); i++ {
		dp[i] = common.Max(dp[i-1], dp[i-2]+newNums[i-1])
	}
	return dp[len(newNums)]
}
