package dynamicprogramming

import "testing"

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000
// https://leetcode-cn.com/problems/longest-palindromic-substring/ medium
func Test_LongestPalindrome(t *testing.T) {
	s := "rtyuuytr"
	t.Log(longestPalindrome(s))
}

// 对于一个字串，假设长度大于2，如果他有两种情况，是回文和不是回文，如果是回文，那么去除首尾两个字符后仍是回文
// 以dp[i][j]表示从字符串从i-j是否为回文，那么dp[i][j] = dp[i+1][j-1] && s[i]==s[j]
// 定义base，即长度小于等于2是，如果长度为1，那么他就是一个回文字串，即dp[i][i] = true
// 如果长度为2，只要这两个字符相等，那么它是一个回文字串，即dp[i][i+1] = s[i]==s[i+1]
// 我们要找的最长回文字串即是dp中为true，j-i最大的(不唯一)
func longestPalindrome(s string) string {
	ret := ""
	dp := make([][]bool, len(s)) // 初始化dp数组
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	// 以字串长度为条件定义base组成状态方程并实现
	for gap := 0; gap < len(s); gap++ {
		for i := 0; i+gap < len(s); i++ {
			j := i + gap  // dp[i][j] 间隔gap 即以子序列长度为标准遍历
			if gap == 0 { // 初始化base
				dp[i][j] = true
			} else if gap == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] && gap+1 > len(ret) {
				ret = s[i : j+1]
			}
		}
	}
	return ret
}
