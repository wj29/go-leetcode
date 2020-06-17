package doublepoint

import (
	"testing"
)

// 给定一个字符串s，反转其中的元音字母(aeiou) easy
// https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/

func Test_ReverseVowels(t *testing.T) {
	str := "aftgee"
	t.Log(reverseVowels(str))
}

// 双端遍历 时间负载度n
func reverseVowels(s string) string {
	if len(s) <= 0 {
		return s
	}
	vowelsMap := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	l := 0
	r := len(s) - 1
	b := []byte(s)
	for {
		for r > 0 && !vowelsMap[string(b[r])] { // l r判断要放在前面，防止数组panic
			r--
		}
		for l < len(s) && !vowelsMap[string(b[l])] {
			l++
		}
		if l < r {
			b[l], b[r] = b[r], b[l]
			l ++
			r --
		} else {
			break
		}
	}
	return string(b)
}
