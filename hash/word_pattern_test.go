package hash

import (
	"strings"
	"testing"
)

// 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律
// 输入: pattern = "abba", str = "dog cat cat dog" 输出: true
// https://leetcode-cn.com/problems/word-pattern/ easy
func Test_WordPattern(t *testing.T) {
	pattern := "abba"
	str := "dog cat cat dog"
	t.Log(pattern, str)
}

func wordPattern(pattern string, str string) bool {
	arr := strings.Split(str, " ")
	if len(pattern) != len(arr) {
		return false
	}
	m := make(map[string]string)
	hasKey := make(map[string]bool)
	for i := 0; i < len(pattern); i++ {
		_, ok := m[string(pattern[i])]
		switch {
		case !ok && !hasKey[arr[i]]:
			m[string(pattern[i])] = arr[i]
			hasKey[arr[i]] = true
		case !ok && hasKey[arr[i]]:
			return false
		case ok && hasKey[arr[i]]:
			if m[string(pattern[i])] != arr[i] {
				return false
			}
		case ok && !hasKey[arr[i]]:
			return false
		}
	}
	return true
}
