package dynamicprogramming

import (
	"testing"
)

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串
// s 可能为空，且只包含从 a-z 的小写字母
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *
// https://leetcode-cn.com/problems/regular-expression-matching hard
func Test_IsMatch(t *testing.T) {
	s := "ab"
	p := ".*"
	t.Log(isMatch(s, p))
}

// 对于字符串字串的问题通常使用动态规划解决
// dp[i][j]表示s的前i个字符与p的前j个字符是否匹配
// 分情况讨论:
// p[j]为字母不是*，如果s[i] != p[j]，dp[i][j]=false，如果s[i] == p[j]，那么dp[i][j] = dp[i-1][j-1]
// p[j]为*，*+前一个字母表示可以取0-n个，即判断dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-2][j-2] ... (从后往前s都字母都一样)
// 变化思路(参考官方题解)(怎么变化)
// 如果s[i] != p[j-1] 比如 abc和abb*此时b(j-1)!=c(i)，dp[i][j] = dp[i][j-2]
// 如果s[i] == p[j-1] 比如 abc和abc*此时c(j-1)==c(i)，dp[i][j] = dp[i-1][j] || dp[i][j-2]
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] != '*' {
				if match(s, p, i, j) {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			} else {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(s, p, i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

func match(s, p string, i, j int) bool {
	if i == 0 {
		return false
	}
	if p[j-1] == '.' {
		return true
	}
	return s[i-1] == p[j-1]
}
