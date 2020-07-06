package dynamicprogramming

import (
	"strings"
	"testing"
)

// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词
// 拆分时可以重复使用字典中的单词，你可以假设字典中没有重复的单词
// https://leetcode-cn.com/problems/word-break/ medium
func Test_WordBreak(t *testing.T) {
	s := "leetgcodeleetg"
	wordDict := []string{"leetg", "code"}
	t.Log(wordBreak(s, wordDict))
	t.Log(wordBreak3(s, wordDict))
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
			if i-len(wordDict[j]) >= 0 && strings.HasSuffix(s[:i], wordDict[j]) {
				dp[i] = dp[i-len(wordDict[j])] || dp[i]
			}
		}
	}
	return dp[len(s)]
}

// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词
// 使得句子中所有的单词都在词典中。返回所有这些可能的句子
// 拆分时可以重复使用字典中的单词，你可以假设字典中没有重复的单词
// https://leetcode-cn.com/problems/word-break-ii/ hard
func Test_WordBreak2(t *testing.T) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	ret := wordBreak3(s, wordDict)
	for _, v := range ret {
		t.Log(v)
	}
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
	dp[0] = []string{""} // 必须有一个初始值才能正确的在后续循环中对dp[i]添加元素(替换)，每一个字典的元素的dp[i]为空，但是会新扩充一个元素为该当前字典
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < len(wordDict); j++ {
			if i-len(wordDict[j]) >= 0 && strings.HasSuffix(s[:i], wordDict[j]) {
				for _, v := range dp[i-len(wordDict[j])] {
					if v == "" {
						v = wordDict[j] // 此处相当于给dp[i]新增了一个元素，扩充了dp[i]
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

// 通过dfs实现，从字符s开始开始搜索
// 一般情况下，在寻求结果集的时候都需要用到递归
// 递归取值的方法一般需要两个数组进行取值
func wordBreak3(s string, wordDict []string) []string {
	if s == "" || len(wordDict) == 0 {
		return []string{}
	}
	// dp这一段代码的作用是面对上述测试用例直接返回false
	// 去掉一样会超时
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < len(wordDict); j++ {
			if i-len(wordDict[j]) >= 0 && strings.HasSuffix(s[:i], wordDict[j]) {
				dp[i] = dp[i-len(wordDict[j])] || dp[i]
			}
		}
	}
	if !dp[len(s)] {
		return []string{}
	}
	ret := make([]string, 0)
	search(s, wordDict, []string{}, &ret)
	return ret
}

func search(s string, wordDict, r []string, ret *[]string) {
	if len(s) == 0 {
		*ret = append(*ret, strings.Join(r, " "))
	}
	for _, word := range wordDict {
		if len(s) < len(word) || !strings.HasPrefix(s, word) {
			continue
		}
		search(s[len(word):], wordDict, append(r, word), ret)
	}
}
