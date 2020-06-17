package hash

import (
	"testing"
)

// 给定两个字符串 s 和 t，判断它们是否是同构的。
// 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身
// https://leetcode-cn.com/problems/isomorphic-strings
func Test_IsIsomorphic(t *testing.T) {
	s := "paper"
	p := "title"
	t.Log(isIsomorphic(s, p))
}

// 通过两个map储存，将数组对应下标当作value
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m1[s[i]] = i
		m2[t[i]] = i
	}
	for i := 0; i < len(s); i++ {
		if m2[t[i]] != m1[s[i]] {
			return false
		}
	}
	return true
}
