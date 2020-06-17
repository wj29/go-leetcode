package hash

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 和谐数组: 数组中最大值与最小值的差值为1
// 给定一个数组，找到最长和谐子序列的长度
// 输入: [1,3,2,2,5,2,3,7] 输出: 5 原因: 最长的和谐数组是：[3,2,2,2,3]
// https://leetcode-cn.com/problems/longest-harmonious-subsequence/description/ easy
func Test_FindLHS(t *testing.T) {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	nums = []int{1, 2, 2, 1}
	t.Log(findLHSByArr(nums))
	t.Log(findLHSByHash(nums))
}

// 对于元素n，他可以和元素n+1组成和谐序列
// 双指针法，遍历数组每一个元素，然后重头开始遍历，时间复杂度N^2
// 对每个元素进行hash，hash后重新遍历，时间复杂度N
func findLHSByArr(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	l := 0
	for i := 0; i < len(nums); i++ {
		count := 0
		flag := false
		for j := 0; j < len(nums); j++ {
			if nums[j] == nums[i]+1 {
				count++
				flag = true
			} else if nums[j] == nums[i] {
				count++
			}
		}
		if flag {
			l = common.Max(l, count)
		}
	}
	return l
}

func findLHSByHash(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	l := 0
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]] ++
		} else {
			m[nums[i]] = 1
		}
	}
	for i := range m {
		if _, ok := m[i+1]; ok {
			l = common.Max(l, m[i]+m[i+1])
		}
	}
	return l
}
