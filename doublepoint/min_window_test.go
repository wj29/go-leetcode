package doublepoint

import (
	"math"
	"testing"
)

// 76. 最小覆盖子串
// https://leetcode-cn.com/problems/minimum-window-substring/
func Test_MinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	c := "ABC"
	s = "a"
	c = "aa"
	t.Log(minWindow(s, c))
}

// 滑动窗口
// 框架解法如下
func minWindow(s string, t string) string {
	left, right := 0, 0 // 双指针滑动
	start, length := 0, math.MaxInt32
	window, need := make(map[byte]int), make(map[byte]int) // 窗口和条件
	for i := range t {
		need[t[i]]++
	}
	for right < len(s) {
		window[s[right]]++ // 扩大窗口
		right++

		if !cover(window, need) { // 不满足则继续
			continue
		}
		for cover(window, need) { // 满足则左侧开始缩小窗口直到不满足
			if right-left < length {
				start, length = left, right-left
			}
			window[s[left]]--
			left++
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

func cover(coverer, covered map[byte]int) bool {
	for k, v := range covered {
		if coverer[k] < v {
			return false
		}
	}
	return true
}
