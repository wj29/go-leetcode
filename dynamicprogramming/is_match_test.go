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
// p[j]为*，*加上前一个字母表示可以取0-n个，取p[j-1]为0次的时候，dp[i][j] = dp[i][j-2]
// 同理，当取p[j-1]为1、2、3次的时候，需要s[i]==p[j-1], s[i]==s[i-1]==p[j-1]，此时dp[i][j] = dp[i-1][j-2]、dp[i-2][j-2]、dp[i-3][j-2]等

// 参考题解: 字母 + 星号的组合在匹配的过程中，本质上只会有两种情况：
// 匹配s末尾的一个字符，将该字符扔掉，而该组合还可以继续进行匹配(此处相当于dp[i][j] = dp[i-1][j]，dp[i-2][j]在上一轮循环中已经判断过)
// 不匹配字符，将该组合扔掉，不再进行匹配。
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
	if i == 0 { // 代表s为“”，在p[j-1]不是‘*’时一定为false
		return false
	}
	if p[j-1] == '.' {
		return true
	}
	return s[i-1] == p[j-1]
}

// 给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
// '?' 可以匹配任何单个字符。
// '*' 可以匹配任意字符串（包括空字符串）。
// https://leetcode-cn.com/problems/wildcard-matching
func isMatch2(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(s) + 1; i ++ {
		for j := 1; j < len(p) + 1; j ++ { // dp[*][0] = false
			if p[j-1] != '*' {
				if match(s, p, i, j) {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				dp[i][j] = dp[i][j-1]
				if i != 0 {
					dp[i][j] = dp[i-1][j] || dp[i][j]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}