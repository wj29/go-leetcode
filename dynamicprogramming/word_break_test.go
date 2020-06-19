package dynamicprogramming

import (
	"testing"
)

// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词
// 拆分时可以重复使用字典中的单词，你可以假设字典中没有重复的单词
// https://leetcode-cn.com/problems/word-break/ medium
func Test_WordBreak(t *testing.T) {
	s := "leetgcodeleetg"
	wordDict := []string{"leetg", "code"}
	t.Log(wordBreak(s, wordDict))
}

// 对于字符s，它可以通过最后一个字串+字典中的任意一个单词构成
// dp[n] = || dp[n-dict[i]](for i 0-len(dict))
func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < len(wordDict); j++ {
			if i-len(wordDict[j]) >= 0 && isLastWordOfString(wordDict[j], s[:i]) {
				dp[i] = dp[i-len(wordDict[j])] || dp[i]
			}
		}
	}
	return dp[len(s)]
}

func isLastWordOfString(word, s string) bool {
	if word == "" || len(word) > len(s) {
		return false
	}
	gap := len(s) - len(word)
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] != s[i+gap] {
			return false
		}
	}
	return true
}

// TODO
// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词
// 使得句子中所有的单词都在词典中。返回所有这些可能的句子
// 拆分时可以重复使用字典中的单词，你可以假设字典中没有重复的单词
// https://leetcode-cn.com/problems/word-break-ii/ hard
func Test_WordBreak2(t *testing.T) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	t.Log(wordBreak2(s, wordDict))
}

// 通过将dp数组转换成二维数组，存储该位置的所有可能
// 面对测试用例 "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
// ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"] 超时
func wordBreak2(s string, wordDict []string) []string {
	if s == "" {
		return []string{}
	}
	if len(wordDict) == 0 {
		return []string{}
	}
	dp := make([][]string, len(s)+1)
	for i := range dp {
		dp[i] = make([]string, 0)
	}
	dp[0] = []string{""}
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < len(wordDict); j++ {
			if i-len(wordDict[j]) >= 0 && isLastWordOfString(wordDict[j], s[:i]) {
				for _, v := range dp[i-len(wordDict[j])] {
					if v == "" {
						v = wordDict[j]
					} else {
						v = v + " " + wordDict[j]
					}
					dp[i] = append(dp[i], v)
				}
			}
		}
	}
	return dp[len(s)]
}

// TODO
// 通过dfs实现，从字符s开始开始搜索
// https://leetcode-cn.com/problems/word-break-ii/solution/dong-tai-gui-hua-qian-zhui-shu-by-scut_dell/
