package arrayandmatrix

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。
// 如果不存在符合条件的连续子数组，返回 0
// https://leetcode-cn.com/problems/minimum-size-subarray-sum
func Test_MinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 7
	t.Log(minSubArrayLen(s, nums))
	t.Log(minSubArrayLen1(s, nums))
}

// 以子数组长度为标准遍历数组，得到大于等于target的最小数组长度，时间复杂度N^2
// 当数组较大时，会花费很长的时间
func minSubArrayLen(s int, nums []int) int {
	ret := 0
	if len(nums) == 0 {
		return ret
	}
	prev := make([]int, len(nums)+1) // 构造前缀和
	prev[0] = 0
	for i := 0; i < len(nums); i++ {
		prev[i+1] = prev[i] + nums[i]
	}
	for gap := 0; gap < len(nums); gap++ {
		for i := 0; i+gap < len(nums); i++ {
			val := 0
			//start := i
			//for start <= i+gap {
			//	val += nums[start]
			//	start++
			//}
			val = prev[i+gap+1] - prev[i]
			if val >= s {
				return gap + 1
			}
		}
	}
	return ret
}

// 以元素结点为结束结点标准遍历数组，当找到第一个符合条件的长度时，后序遍历只需要目前长度的累加和即可
// 时间复杂度N^2

// 通过双指针和滑动窗口解决
// 当start和end指针从数组开始移动，end指针向右移动直到第一次满足>=target记录下end-start
// 由于是正整数数组，此时end指针继续向右移动一定满足>=target，此时将start指针向右移动，直到第一次不能满足>=target
// 继续移动右指针，重复上述过程直到end达到数组末尾，start到end和再一次不能满足>=target为止
func minSubArrayLen1(s int, nums []int) int {
	ret := 0
	if len(nums) == 0 {
		return ret
	}
	start, end := 0, 0               // 定义双指针
	prev := make([]int, len(nums)+1) // 构造前缀和
	prev[0] = 0
	for i := 0; i < len(nums); i++ {
		prev[i+1] = prev[i] + nums[i]
	}
	for end < len(nums) && start <= end {
		tmpValue := 0
		//for i := start; i <= end; i++ {
		//	tmpValue += nums[i] // 此段代码使用前缀和实现，避免循环累加，从160ms优化到8ms
		//}
		tmpValue = prev[end+1] - prev[start]
		if tmpValue >= s {
			if ret == 0 {
				ret = end - start + 1
			} else {
				ret = common.Min(ret, end-start+1)
			}
			start++
		} else {
			end++
		}
	}
	return ret
}
