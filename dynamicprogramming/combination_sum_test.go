package dynamicprogramming

import (
	"sort"
	"testing"
)

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合
// candidates 中的数字可以无限制重复被选取
// https://leetcode-cn.com/problems/combination-sum medium
func Test_CombinationSum(t *testing.T) {
	candidates := []int{3, 2, 7}
	target := 18
	t.Log(combinationSum(candidates, target))
}

// 递归，类似于二叉树的根结点到叶子结点和为target的路径 /tree/path_sum_test.go
// 回溯法
func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	for i := range ret {
		ret[i] = make([]int, 0)
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	path(candidates, target, 0, []int{}, &ret)
	return ret
}

func path(candidates []int, target int, start int, prev []int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(prev)) // 不能直接append，会导致错误
		copy(tmp, prev)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ { // start作用是用过的数据不可以重复使用，避免重复解
		if target < candidates[i] {
			return
		}
		path(candidates, target-candidates[i], i, append(prev, candidates[i]), res)
	}
}

// 根据兑换零钱的解法，我们知道对于candidates中的元素，我们可以选择或者不选择
// dp[i][]表示对于第k个元素能组合成target的种类数
func combinationSum1(candidates []int, target int) [][]int {
	dp := make([][]int, target+1)
	for i := range dp {
		dp[i] = make([]int, 0)
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})
	for i := 0; i < len(candidates); i++ {
		// TODO: how to implement?
	}
	return [][]int{}
}
