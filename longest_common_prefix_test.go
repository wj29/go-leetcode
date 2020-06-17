package leetcode

import (
	"testing"
)

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""
// https://leetcode-cn.com/problems/longest-common-prefix/  easy
func Test_LongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	t.Log(longestCommonPrefix(strs))
}

// 通过找到最小的字符串长度，遍历每一个元素0-最小长度的字符是否相等
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	minLength := len(strs[0])
	for _, v := range strs {
		if len(v) < minLength {
			minLength = len(v)
		}
	}
	index := 0
loop:
	for i := 0; i < minLength; i++ {
		tmp := string(strs[0][i])
		for _, v := range strs {
			if string(v[i]) != tmp {
				break loop
			}
		}
		index++
	}
	return strs[0][:index]
}
