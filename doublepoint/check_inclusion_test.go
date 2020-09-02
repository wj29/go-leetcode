package doublepoint

import (
	"reflect"
	"testing"
)

// 567. 字符串的排列
// https://leetcode-cn.com/problems/permutation-in-string/
func Test_CheckInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	t.Log(checkInclusion(s1, s2))
}

// 滑动窗口，判断window和need是否相等即可
// 由于是子串，不需要使用left right双指针移动，因为每次步长相同，且选择的字符串长度为s1长度
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	if len(s1) > len(s2) {
		return false
	}
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
		if i == len(s1)-1 {
			continue
		}
		window[s2[i]]++
	}
	for i := 0; i+len(s1) <= len(s2); i++ {
		if i > 0 {
			if window[s2[i-1]] == 1 {
				delete(window, s2[i-1])
			} else {
				window[s2[i-1]]--
			}
		}
		window[s2[i+len(s1)-1]]++
		if reflect.DeepEqual(window, need) {
			return true
		}
	}
	return false
}
