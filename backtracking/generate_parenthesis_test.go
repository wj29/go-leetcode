package backtracking

import (
	"testing"
)

// 数字n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// https://leetcode-cn.com/problems/generate-parentheses/
func Test_GenerateParenthesis(t *testing.T) {
	n := 3
	t.Log(generateParenthesis(n))
}

// 回溯法，每一种遍历都去尝试，)的个数不能大于(的个数
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	ret := make([]string, 0)
	backTraceForG(n, []string{}, &ret, 0, 0)
	return ret
}

func backTraceForG(n int, prev []string, ret *[]string, aCnt, dCnt int) {
	if len(prev) != 0 && len(prev[0]) == 2*n {
		tmp := make([]string, len(prev))
		copy(tmp, prev)
		*ret = append(*ret, tmp...)
		return
	}
	if aCnt+1 <= n {
		if len(prev) == 0 {
			prev = append(prev, "(")
		} else {
			for i := range prev {
				prev[i] += "("
			}
		}
		backTraceForG(n, prev, ret, aCnt+1, dCnt)
		for i := range prev {
			prev[i] = prev[i][:len(prev[i])-1]
		}
	}
	if dCnt+1 <= aCnt {
		if len(prev) == 0 {
			prev = append(prev, "(")
		} else {
			for i := range prev {
				prev[i] += ")"
			}
		}
		backTraceForG(n, prev, ret, aCnt, dCnt+1)
		for i := range prev {
			prev[i] = prev[i][:len(prev[i])-1]
		}
	}
	return
}
