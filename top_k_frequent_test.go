package leetcode

import (
	"testing"
)

// 给定一个非空整数数组，返回出现频率最高的前K个元素 medium
// https://leetcode-cn.com/problems/top-k-frequent-elements/description/
// https://leetcode-cn.com/problems/sort-characters-by-frequency/description/ 同解

func Test_TopKFrequent(t *testing.T) {
	nums := []int{1}
	k := 1
	t.Log(topKFrequent(nums, k))
}

// 通过map的唯一key将元素和出现次数作为key和value存入map，利用二维数组循环map，出现的次数作为数组下标，值为出现的元素
// 倒序遍历得到频率为k的元素，时间复杂度为2N+K即N
func topKFrequent(nums []int, k int) []int {
	if len(nums) <= 0 || k < 1 || k > len(nums) {
		panic("invalid")
	}
	bucket := make([][]int, len(nums)+1)
	m := make(map[int]int) // key为数组元素，value为出现次数
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]] += 1
		} else {
			m[nums[i]] = 1
		}
	}
	for i, v := range m {
		bucket[v] = append(bucket[v], i)
	}
	r := make([]int, 0)
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, v := range bucket[i] {
			r = append(r, v)
			if len(r) >= k {
				return r
			}
		}
	}
	return r
}
