package arrayandmatrix

import "testing"

// 给定两个数组，编写一个函数来计算它们的交集
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致
// 我们可以不考虑输出结果的顺序
// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/ easy

func Test_Intersect(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	t.Log(intersect(nums1, nums2))
}

// 通过map实现计数即可
func intersect(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	m := make(map[int]int)
	for i := range nums1 {
		if _, ok := m[nums1[i]]; !ok {
			m[nums1[i]] = 1
		} else {
			m[nums1[i]] ++
		}
	}
	for i := range nums2 {
		if _, ok := m[nums2[i]]; ok {
			ret = append(ret, nums2[i])
			m[nums2[i]]--
			if m[nums2[i]] == 0 {
				delete(m, nums2[i])
			}
		}
	}
	return ret
}
