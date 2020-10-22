package greedy

import (
	"testing"
)

// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段
// 返回一个表示每个字符串片段的长度的列表
// 输入: S = "ababcbacadefegdehijhklij"
// 输出: [9,7,8] 划分结果为 "ababcbaca", "defegde", "hijhklij"
// 每个字母最多出现在一个片段中
// 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少
// 763. 划分字母区间
// https://leetcode-cn.com/problems/partition-labels medium
func Test_PartitionLabels(t *testing.T) {
	s := "ababcbacadefegdehijhklij"
	t.Log(partitionLabels(s))
}

func partitionLabels(S string) []int {
	m := make(map[string]int)
	for i := 0; i < len(S); i++ {
		m[string(S[i])] = i
	}
	begin := 0
	current := 0
	last := -1
	ret := make([]int, 0)
	for current < len(S) {
		last = max(m[string(S[current])], last)
		if current < last {
			current++
		} else if current == last {
			ret = append(ret, last-begin+1)
			current++
			begin = current
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
