package arrayandmatrix

import (
	"testing"
)

// 集合S包含从1到n的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。
// 给定一个数组nums(无序)代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回
// 输入: nums = [1,2,2,4]
// 输出: [2,3]
// https://leetcode-cn.com/problems/set-mismatch  easy
func Test_FindErrorNums(t *testing.T) {
	nums := []int{1, 2, 2, 4}
	t.Log(findErrorNums(nums))
}

// 排序后遍历，时间复杂度为排序复杂度+N
// 使用map，值等于2就是多的，等于0的就是缺失的 时间复杂度N 空间复杂度N
// 使用额外的数组和map可以达到相同效果，对应下标存入对应值 时间复杂度N 空间复杂度N
func findErrorNums(nums []int) []int {
	dup, missing := -1, -1
	m := make(map[int]int)

	for i := range nums {
		m[i+1] = 0
	}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		}
	}
	for k, v := range m {
		if v == 2 {
			dup = k
		}
		if v == 0 {
			missing = k
		}
	}
	return []int{dup, missing}
}
