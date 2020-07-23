package backtracking

import (
	"sort"
	"testing"
)

// 给定一个 没有重复 数字的序列，返回其所有可能的全排列
// https://leetcode-cn.com/problems/permutations/

// 给定一个可包含重复数字的序列，返回所有不重复的全排列
// https://leetcode-cn.com/problems/permutations-ii/
func Test_Permute(t *testing.T) {
	nums := []int{2, 1, 3}
	t.Log(insertEleToSlice(nums, 4))
	t.Log(permuteByDP(nums))
	t.Log(permute(nums))
	nums = []int{1, 1, 2}
	t.Log(permuteUnique(nums))
}

// 回溯法解决(模板)
// def backtrack(路径, 选择列表):
//    if 满足结束条件:
//        result.add(路径)
//        return
//
//    for 选择 in 选择列表:
//        做选择
//        backtrack(路径, 选择列表)
//        撤销选择
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	backTrackForPermute(nums, []int{}, &ret)
	return ret
}

func backTrackForPermute(nums []int, prev []int, ret *[][]int) {
	/*
		if len(nums) == 0 { // 使用len(nums)==0判断在for循环中nums不方便传递，使用另一个数组表示已经被选择的数据即prev
			tmp := make([]int, len(prev))
			copy(tmp, prev)
			*ret = append(*ret, tmp)
		}
	*/
	if len(nums) == len(prev) {
		tmp := make([]int, len(prev))
		copy(tmp, prev)
		*ret = append(*ret, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if isInSlice(prev, nums[i]) { // 已经被选择过
			continue
		}
		// 保证每一次循环中的prev是一样的
		// prev = append(prev, nums[i])
		// backTrackForPermute(nums, prev, ret)
		// prev = prev[:len(prev)-1]

		backTrackForPermute(nums, append(prev, nums[i]), ret)
	}
}

func isInSlice(s []int, ele int) bool {
	for _, v := range s {
		if v == ele {
			return true
		}
	}
	return false
}

// 动态规划
// 用dp[i][j]表示到第i个元素的j个全排列，它的值是dp[i-1]的值分别加上将nums[i]插入dp[i-1]中
// 即 dp[i][j] = dp[i-1][j](for j in 0 - i-1), num[i]插入
func permuteByDP(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	dp := make([][][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([][]int, 0)
	}
	dp[1] = append(dp[1], []int{nums[0]})
	for i := 2; i < len(dp); i++ {
		for _, s := range dp[i-1] {
			r := insertEleToSlice(s, nums[i-1])
			dp[i] = append(dp[i], r...)
		}
	}
	return dp[len(nums)]
}

func insertEleToSlice(s []int, ele int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < len(s)+1; i++ {
		tmp := make([]int, 0)
		copy(tmp, s)
		tmp = append(tmp, s[:i]...)
		tmp = append(tmp, ele)
		tmp = append(tmp, s[i:]...)
		ret = append(ret, tmp)
	}
	return ret
}

// 此题的难点是如何去重
// 去重方式和combination_sum_test.go一样
func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)
	for i := range ret {
		ret[i] = make([]int, 0)
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	backTrackForPermuteUnique(nums, []int{}, &ret)
	return ret
}

func backTrackForPermuteUnique(nums []int, prev []int, ret *[][]int) {
	
}
