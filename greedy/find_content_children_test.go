package greedy

import (
	"testing"
)

// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干
// 对每个孩子 i ，都有一个胃口值 gi ，这是能让孩子们满足胃口的饼干的最小尺寸
// 并且每块饼干 j ，都有一个尺寸 sj 。如果 sj >= gi ，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足
// 你的目标是尽可能满足越多数量的孩子，并输出这个最大数值
// eg 输入 [1,2,3], [1,1] 输出: 1  输入 [1,2], [1,2,3] 输出: 2 easy
// https://leetcode-cn.com/problems/assign-cookies/description/

func Test_FindContentChildren(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	t.Log(findContentChildren(g, s))
}

// 贪心算法，使用最小尺寸的饼干满足最小胃口的孩子
func findContentChildren(g []int, s []int) int {
	// 先排序
	qSort(g, 0, len(g)-1)
	qSort(s, 0, len(s)-1)
	child := 0
	cookie := 0
	for child < len(g) && cookie < len(s) {
		if s[cookie] >= g[child] {
			cookie++
			child++
		} else {
			cookie++
		}
	}
	return child
}

func qSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := nums[start]
	l := start
	r := end
	for {
		for nums[r] >= pivot && r > start {
			r--
		}
		for nums[l] <= pivot && l < end {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		} else {
			break
		}
	}
	nums[start], nums[r] = nums[r], nums[start]
	qSort(nums, start, r)
	qSort(nums, r+1, end)
}