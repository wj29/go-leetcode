package string

import "testing"

// 28. 实现 strStr()
// https://leetcode-cn.com/problems/implement-strstr/
func Test_StrStr(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	t.Log(strStr(haystack, needle))
}

// KMP算法
// 构建next数组 https://blog.csdn.net/yearn520/article/details/6729426
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	// 构建next数组
	next := make([]int, len(needle))
	for i := 1; i < len(needle); i++ {
		pre := next[i-1]
		for pre != 0 && needle[i] != needle[pre] { // 找到更小的相似程度子串
			pre = next[pre-1]
		}
		if needle[i] == needle[pre] { // 相似程度+1
			next[i] = pre + 1
		}
	}

	i, j := 0, 0
	for j < len(haystack) {
		if haystack[j] == needle[i] {
			i++
			j++
		}
		if i == len(needle) {
			return j - i
		}
		if j < len(haystack) && haystack[j] != needle[i] {
			if i == 0 {
				j++
			} else {
				i = next[i-1]
			}
		}
	}
	return -1
}
