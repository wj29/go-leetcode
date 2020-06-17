package arrayandmatrix

import "testing"

// 给定两个数组，编写一个函数来计算它们的交集
// 输出结果中的每个元素一定是唯一的
// 我们可以不考虑输出结果的顺序
// https://leetcode-cn.com/problems/intersection-of-two-arrays/ easy
func Test_Intersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	t.Log(intersection(nums1, nums2))
}

// 两个数组的交集，将一个数组存入map，用另一个数组去匹配
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	for _, v := range nums1 {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	ret := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			ret = append(ret, v)
			delete(m, v)
		}
	}
	return ret
}
