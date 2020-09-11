package backtracking

import (
	"github.com/wujiang1994/go-leetcode/common"
	"sort"
	"testing"
)

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合
// 所有数字（包括 target）都是正整数
// 解集不能包含重复的组合

// candidates 中的数字可以无限制重复被选取
// https://leetcode-cn.com/problems/combination-sum medium

// candidates 中的每个数字在每个组合中只能使用一次。
// https://leetcode-cn.com/problems/combination-sum-ii/
func Test_CombinationSum(t *testing.T) {
	candidates := []int{3, 2, 6, 7}
	target := 7
	t.Log(combinationSum(candidates, target, false))
	t.Log(combinationSum(candidates, target, true))
	k, n := 3, 7
	k, n = 3, 9
	t.Log(combinationSum3(k, n))
}

// 递归，类似于二叉树的根结点到叶子结点和为target的路径 /tree/path_sum_test.go
// 回溯法
func combinationSum(candidates []int, target int, isOnceOnly bool) [][]int {
	ret := make([][]int, 0)
	for i := range ret {
		ret[i] = make([]int, 0)
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	path(candidates, target, 0, []int{}, &ret, isOnceOnly)
	return ret
}

func path(candidates []int, target int, start int, prev []int, res *[][]int, isOnceOnly bool) {
	if target == 0 {
		tmp := make([]int, len(prev)) // 不能直接append，会导致错误
		copy(tmp, prev)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ { // start作用是用过的数据不可以重复使用，避免重复解
		if i > start && isOnceOnly && candidates[i] == candidates[i-1] {
			continue
		}
		if target < candidates[i] {
			return
		}
		index := i
		if isOnceOnly {
			index += 1
		}
		path(candidates, target-candidates[i], index, append(prev, candidates[i]), res, isOnceOnly) // 此处prev未变
	}
}

// 216. 组合总和 III
// https://leetcode-cn.com/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	backTraceForCombinationSum3(&ret, []int{}, 1, n, k)
	return ret
}

func backTraceForCombinationSum3(ret *[][]int, prev []int, start, n, k int) {
	if len(prev) == k && common.Sum(prev) == n {
		tmp := make([]int, k)
		copy(tmp, prev)
		*ret = append(*ret, tmp)
		return
	}
	for i := start; i < 10; i++ {
		if len(prev) >= k {
			continue
		}
		backTraceForCombinationSum3(ret, append(prev, i), i+1, n, k)
	}
}
