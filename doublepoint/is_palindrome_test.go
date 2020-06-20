package doublepoint

import (
	"strings"
	"testing"
)

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// 说明：本题中，我们将空字符串定义为有效的回文串。
// https://leetcode-cn.com/problems/valid-palindrome/ easy
func Test_IsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	s = "race a car"
	t.Log(isPalindrome(s))
}

// 将字符串只保留大小写字母和数字，字母可都改成小写
// 双指针分别从头和尾向中间逼近，判断是否相等
func isPalindrome(s string) bool {
	t := ""
	for _, v := range s {
		tmp := strings.ToLower(string(v))
		if (tmp <= "9" && tmp >= "0") || (tmp >= "a" && tmp <= "z") {
			t += tmp
		}
	}
	if t == "" {
		return true
	}
	l, r := 0, len(t)-1
	for l < r {
		if t[l] != t[r] {
			return false
		}
		l++
		r--
	}
	return true
}
