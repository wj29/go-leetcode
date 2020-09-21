package arrayandmatrix

import (
	"sort"
	"testing"
)

// 1589. 所有排列中的最大和
// https://leetcode-cn.com/problems/maximum-sum-obtained-of-any-permutation/
func Test_MaxSumRangeQuery(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	requests := [][]int{{1, 3}, {0, 1}}
	t.Log(maxSumRangeQuery(nums, requests))
}

// 给定一个数组nums，差分数组一般用于对其中子数组nums[i:j]统一操作
// d表示nums的差分数组，那么d[0] = nums[0], d[i] = nums[i]-nums[i-1]
// 对区间i-j操作+count，即使用差分数组操作为d[i]+=count, d[j+1]-=count(j+1存在)
// 通过变化后的差分数组，求变化后的原数组，的nums[i] = nums[i-1]+d[i]，其中a[0]=d[0]

// requests中出现最多的下标应该选择nums中最大的数，贪心思想
func maxSumRangeQuery(nums []int, requests [][]int) int {
	sort.Ints(nums)
	ret := 0
	d := make([]int, len(nums)) // 所有下标出现0次
	for _, request := range requests {
		d[request[0]]++
		if request[1]+1 >= len(nums) {
			continue
		}
		d[request[1]+1]--
	}
	// 利用差分数组求变化后原数组 count[i] = count[i-1] + d[i]
	count := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			count[i] = d[i]
		} else {
			count[i] = count[i-1] + d[i]
		}
	}
	sort.Ints(count)

	for i := len(nums) - 1; i >= 0; i-- {
		if count[i] == 0 {
			break
		}
		ret += count[i] * nums[i]
	}
	return ret % (1e9 + 7)
}
