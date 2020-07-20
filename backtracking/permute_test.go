package backtracking

import "testing"

// 给定一个 没有重复 数字的序列，返回其所有可能的全排列
// https://leetcode-cn.com/problems/permutations/
func Test_Permute(t *testing.T) {

}

// 回溯法解决
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	pathForPermute(nums, []int{}, &ret)
	return ret
}

func pathForPermute(nums []int, prev []int, ret *[][]int) {

}
