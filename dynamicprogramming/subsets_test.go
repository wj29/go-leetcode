package dynamicprogramming

import (
	"testing"
)

// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// https://leetcode-cn.com/problems/subsets/ medium
func Test_Subsets(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	ret := subsets(nums)
	for _, v := range ret {
		t.Log(v)
		t.Log("------")
	}
	nums = []int{1, 2, 2}
	ret = subsetsWithDup(nums)
	for _, v := range ret {
		t.Log(v)
		t.Log("------")
	}
}

// 对于数组nums[n]它的子集由nums[n-1]的子集分别加上第n个数组成
// dp[i]表示nums[i]的所有子集，那么对于dp[i]来说，它是通过dp[i-1]所有子集循环加上nums[i]这个数再加上dp[i-1]得到
// 我们得到dp方程 dp[i] = dp[i-1] + append(dp[i-1][j](for j in 0-length), nums[i])
// 注意golang中切片是引用类型，使用内置copy函数得到新的内存地址
func subsets(nums []int) [][]int {
	dp := make([][][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([][]int, 0)
	}
	dp[0] = [][]int{{}}
	for i := 1; i <= len(nums); i++ {
		dp[i] = append(dp[i], dp[i-1]...)
		tmp := make([][]int, len(dp[i-1]))
		copy(tmp, dp[i-1])
		for _, v := range tmp {
			t := make([]int, len(v))
			copy(t, v)
			t = append(t, nums[i-1])
			dp[i] = append(dp[i], t)
		}
	}
	return dp[len(nums)]
}

// 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// https://leetcode-cn.com/problems/subsets-ii/

// 对于含有重复元素的数组分情况讨论，当元素nums[i]未出现过时按照1的方法解决
// 当元素nums[i]出现过时，那么在循环的时候只需要对于该数组里面重复的个数等于元素nums[i]出现的次数-1的进行append，其余的数组全部舍弃(全部已经出现过)
func subsetsWithDup(nums []int) [][]int {
	m := make(map[int]int)
	dp := make([][][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([][]int, 0)
	}
	dp[0] = [][]int{{}}
	for i := 1; i <= len(nums); i++ {
		dp[i] = append(dp[i], dp[i-1]...)
		_, ok := m[nums[i-1]]
		if !ok {
			m[nums[i-1]] = 1
		} else {
			m[nums[i-1]]++
		}
		tmp := make([][]int, len(dp[i-1]))
		copy(tmp, dp[i-1])
		for _, v := range tmp {
			if ok && findCount(v, nums[i-1]) < m[nums[i-1]]-1 {
				continue
			}
			t := make([]int, len(v))
			copy(t, v)
			t = append(t, nums[i-1])
			dp[i] = append(dp[i], t)
		}
	}
	return dp[len(nums)]
}

func findCount(nums []int, ele int) int {
	ret := 0
	for _, v := range nums {
		if v == ele {
			ret++
		}
	}
	return ret
}
