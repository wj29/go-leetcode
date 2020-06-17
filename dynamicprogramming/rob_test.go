package dynamicprogramming

import (
	"testing"
)

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额
// https://leetcode-cn.com/problems/house-robber  easy

// 这个沿街的房屋是环形的，意味着偷了第一家就不能偷最后一家
// https://leetcode-cn.com/problems/house-robber-ii/ medium
func Test_Rob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	nums = []int{2, 7, 9, 3, 1}
	t.Log(rob(nums))
	t.Log(rob2(nums))
}

// 由于只能最少间隔一家，假若f(n-1)的最大值不包括nums[n-1]，那么f(n)的值一定是f(n-1)+nums[i]，中间间隔了nums[n-1]
// 假若f(n-1)取最大值包括了nums[n-1]，那么f(n)的值不能偷nums[n]，会触发报警，那么f(n)应该是max(f(n-2)+nums[n]，f(n-1))中较大的
// 即f(n-2)加上偷nums[n]与f(n-1)不能偷nums[n]中较大者
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	ret := make([]*dpValue, len(nums)+1)
	for i := range ret {
		ret[i] = new(dpValue)
	}
	ret[0] = &dpValue{
		val:  0,
		flag: false,
	}
	ret[1] = &dpValue{
		val:  nums[0],
		flag: true,
	}
	for i := 2; i < len(nums)+1; i++ {
		if ret[i-1].flag {
			if ret[i-1].val > ret[i-2].val+nums[i-1] {
				ret[i].val = ret[i-1].val
				ret[i].flag = false
			} else {
				ret[i].val = ret[i-2].val + nums[i-1]
				ret[i].flag = true
			}
		} else {
			ret[i].val = ret[i-1].val + nums[i-1]
			ret[i].flag = true
		}
	}
	return ret[len(nums)].val
}

type dpValue struct {
	val  int
	flag bool
}

// 这里我们简化掉上一个题解中flag参数，因为无论偷没偷上一家，都不会影响到f(n)最后的取值
// 假若f(n-1)的最大值不包括nums[n-1]，那么f(n)的值一定是f(n-1)+nums[i]，此时f(n-1)=f(n-2)
// 假若f(n-1)取最大值包括了nums[n-1]，那么f(n)的值不能偷nums[n]，那么f(n) = max(f(n-1), f(n-2)+nums[n])
// 所以最终的动态方程为 dp[n] = max(dp[n-1), dp[n-2]+nums[n])
// 此处变成了环形房屋，那么对于任意一家来言，只有偷与不偷
// 以第一家拆分，环形被拆分成两个队列，偷0到n-1家和偷1到n家，取两者最大值
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp1, dp2 := make([]int, len(nums)-1), make([]int, len(nums)-1)
	nums1, nums2 := nums[:len(nums)-1], nums[1:]
	dp1[0], dp2[0] = nums1[0], nums2[0] // 初始化base
	dp1[1], dp2[1] = max(nums1[0], nums1[1]), max(nums2[0], nums2[1])
	for i := 2; i < len(nums)-1; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums1[i])
		
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums2[i])
	}
	return max(dp1[len(nums)-2], dp2[len(nums)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
