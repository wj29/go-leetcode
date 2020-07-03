package dynamicprogramming

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度
// https://leetcode-cn.com/problems/longest-common-subsequence/  medium
func Test_LongestCommonSubsequence(t *testing.T) {
	text1 := "abc"
	text2 := "av"
	t.Log(longestCommonSubsequence(text1, text2))
}

// 一般寻找子序列的题目，都是用动态规划去解决
// 对于text1 text2两个字符串一般可以建立一个二维dp table，通过0-n个字符的字符串去寻找规律
// 题解:
// https://leetcode-cn.com/problems/longest-common-subsequence/solution/dong-tai-gui-hua-zhi-zui-chang-gong-gong-zi-xu-lie/
func longestCommonSubsequence(text1 string, text2 string) int {
	ret := make([][]int, len(text1)+1)
	for i := range ret {
		ret[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				ret[i][j] = ret[i-1][j-1] + 1
			} else {
				ret[i][j] = common.Max(ret[i-1][j], ret[i][j-1])
			}
		}
	}
	return ret[len(text1)][len(text2)]
}
