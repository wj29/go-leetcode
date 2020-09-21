package hash

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 给定一个未排序的整数数组，找出最长连续序列的长度
// 输入 [100, 4, 200, 1, 3, 2] 输出 4 最长连续序列[1, 2, 3, 4]
// 输入 [0, 0, 0, 0] 输出 1 最长序列 [0]
// 要求时间复杂度为N
// https://leetcode-cn.com/problems/longest-consecutive-sequence/description/ hard
func Test_LongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	//nums = []int{0, 0, 0, 0}
	//nums = []int{0, 0, 1}
	t.Log(longestConsecutive(nums))
}

// 排序，再遍历，时间复杂度为排序的时间复杂度
// 遍历对每个元素进行hash，记录每个元素的个数，hash查找的时间复杂度为1，遍历时间复杂度为N
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if !m[v] {
			m[v] = true
		}
	}
	l := 0
	for i := range m {
		count := 0
		for {
			if m[i] {
				count += 1
				i += 1
			} else {
				break
			}
		}
		l = common.Max(l, count)
	}
	return l
}