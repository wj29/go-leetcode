package dynamicprogramming

import "testing"

// 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等
// https://leetcode-cn.com/problems/partition-equal-subset-sum/  medium
func Test_CanPartition(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	nums = []int{1, 2, 3, 5}
	t.Log(canPartition(nums))
}

// 0-1背包问题
// 对于一个容量N的背包，要用这个背包装下物品的价值最大，这些物品有两个属性：体积 w 和价值 v
// 定义dp[i][j] 表示从第0-i件物品中选出总体积不超过j的最大价值
// 那么对于第i件物品，它存在两种属性，加入和不加入
// 第i件物品没添加到背包，总体积不超过j的前i件物品的最大价值就是总体积不超过j的前i-1件物品的最大价值，dp[i][j] = dp[i-1][j]
// 第i件物品添加到背包，dp[i][j] = dp[i-1][j-w]+v
// dp[i][j] = max(dp[i-1][j], dp[i-1][j-w]+v)

// 对于拆分数组成两个子集，即在nums中找到一些元素使得其和等于sum/2
// dp[i][j]表示从数组的0-i元素中选择元素使得其和等于j
// dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum, dp := 0, make([][]bool, len(nums))
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	for i := range dp {
		dp[i] = make([]bool, sum/2+1)
	}
	dp[0][0] = false // nums为正整数
	if nums[0] <= sum/2 { // 从0-0中选择元素使其和等于j，它只能等于它自己，其他默认false
		dp[0][nums[0]] = true
	} else {
		return false
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= sum/2; j++ {
			dp[i][j] = dp[i-1][j]
			if dp[i][j] { // 直接满足
				continue
			}
			if j > nums[i] { // 必须取第i个元素
				dp[i][j] = dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][sum/2]
}
