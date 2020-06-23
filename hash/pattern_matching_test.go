package hash

import (
	"testing"
)

// https://leetcode-cn.com/problems/pattern-matching-lcci/ medium
func Test_PatternMatching(t *testing.T) {
	pattern := "abba"
	value := "dogcatcatdog"
	value = "dogcatcatfish"
	pattern = "aaaa"
	value = "dogcatcatdog"
	pattern = "abba"
	value = "dogdogdogdog"
	pattern = "a"
	value = ""
	pattern = "ab"
	t.Log(patternMatching(pattern, value))
}

// 回溯法尝试所有的可能
func patternMatching(pattern string, value string) bool {
	m := make(map[string]*mVal) // 保存ab对应的字符串
	zSet := make(map[string]struct{}) // 保存所有ab对应的val
	return match(pattern, value, m, zSet)
}

func match(pattern, value string, m map[string]*mVal, zSet map[string]struct{}) bool {
	if pattern == "" {
		return value == ""
	}
	first := pattern[0]                // 不是a就是b
	if v, ok := m[string(first)]; ok { // 存在就比较当前并比较剩下的 刚开始未考虑到a或者b代表""
		if len(value) < len(v.ele) || value[:len(v.ele)] != v.ele {
			return false
		}
		return match(pattern[1:], value[len(v.ele):], m, zSet) // 递归剩下的子字符串
	}
	// 不存在则将当前对应的kv存入map
loop:                                  // 遍历所有情况
	for i := -1; i < len(value); i++ { // ab可以对应到""空字符串
		// a和b不能代表相同的字符串
		for k := range zSet {
			if k == value[:i+1] {
				continue loop
			}
		}
		zSet[value[:i+1]] = struct{}{}
		m[string(first)] = &mVal{ele: value[:i+1]}
		// map的ab有值时，递归找到剩下所有的情况
		if match(pattern[1:], value[i+1:], m, zSet) {
			return true
		}
		delete(m, string(first))
		delete(zSet, value[:i+1])
	}
	return false
}

// 本想使用结构体兼容""，但是在a和b不能同时为一个字符串("")时无法使用
// 可以使用，不想修改
type mVal struct {
	ele string
}
