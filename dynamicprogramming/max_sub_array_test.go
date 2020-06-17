package dynamicprogramming

import "testing"

// 给定一个整数数组nums找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和
// 输入: [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6 [4,-1,2,1]
// https://leetcode-cn.com/problems/maximum-subarray/description/  easy
func Test_MaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(nums))
}

// 暴力法: 遍历找到所有连续子数组，得到最大和 时间复杂度2^N
// 动态规划: 我们以f(i)代表以下标i元素结尾的最大和，那么最后的结构应该是max(f(i))，所以我们只要知道f(i)的值即可
// 举个例子对于abcde中以c结尾的最大值f(c)，不论它本身是c还是bc或者abc，那么f(d)的值一定是c+d，bc+d，abc+d，d中最大值
// 前三项的最大值一定是f(c)+d，那么只需要比较f(c)+d与d的大小即可
// 即f(i)必定是f(i-1) + nums[i]或者不加上nums[i](即nums[i]单独成为开始)，可得到状态方程
// f(i) = max(f(i-1)+nums[i], nums[i])
// 佩服
func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		panic("invalid length")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1] // 累加，用nums[i]的代替f(i)的值
		}
		if nums[i] > max { // 取最大的f(i)
			max = nums[i]
		}
	}
	return max
}
