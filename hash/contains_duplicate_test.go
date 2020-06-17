package hash

import "testing"

// 给定一个整数数组，判断是否存在重复元素
// https://leetcode-cn.com/problems/contains-duplicate/description/ easy
func Test_ContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	t.Log(containsDuplicate(nums))
}

// map实现
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return false
}
